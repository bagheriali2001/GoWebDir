package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/bagheriali2001/GoWebDir/fileHandler"
	"github.com/bagheriali2001/GoWebDir/helpers"
	"github.com/bagheriali2001/GoWebDir/templates"
	"github.com/bagheriali2001/GoWebDir/types"
)

type HandlerWrapper types.HandlerWrapper

func (h *HandlerWrapper) Handler(w http.ResponseWriter, r *http.Request) {
	helpers.HTTPLog(r, http.StatusOK)

	rootPath := h.RootPath

	currentPath := "." + r.URL.Path
	filePath := filepath.Join(rootPath, currentPath)
	if filePath[len(filePath)-1] == '.' {
		filePath = filePath + "/"
	}

	// check if the path is a file
	if filepath.Ext(filePath) != "" {
		http.ServeFile(w, r, filePath)
	} else {
		// list files in directory
		files, err := fileHandler.ListDir(filePath, h.ShowHiddenFiles, h.ShowHiddenFolders)
		if err != nil {
			http.Error(w, "Unable to list directory", http.StatusInternalServerError)
			return
		}

		// if currentPath is not "." then add a escape as first file to files list
		if currentPath != "./" {
			files = append([]types.File{
				{
					Name:    "..",
					Size:    types.Size{},
					Created: "",
					Type:    "Directory",
				},
			}, files...)
		}

		data := types.PageData{
			Directory: types.Directory{
				Name:  currentPath[1:],
				Value: currentPath,
			},
			Files:  files,
			IsRoot: currentPath == "./",
		}

		responseType := "html"
		// check the url query if responseType is set to json
		if r.URL.Query().Get("responseType") == "json" {
			responseType = "json"
		}

		if responseType == "html" {
			// Get minified writer
			mw := templates.GetMinifiedWriter("text/html", w)
			htmlTemplate := templates.GetHTMLTemplate()

			// get the result of htmlTemplate and write it to mw
			var buf bytes.Buffer

			htmlTemplate.Execute(&buf, data)

			bufByte := buf.Bytes()

			mw.Write(bufByte)

			if err := mw.Close(); err != nil {
				http.Error(w, "Unable to minify HTML", http.StatusInternalServerError)
				return
			}

		} else if responseType == "json" {
			// return json response
			jsonData, err := json.Marshal(data)
			if err != nil {
				http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}
