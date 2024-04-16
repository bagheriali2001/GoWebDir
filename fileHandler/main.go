package fileHandler

import (
	"fmt"
	"os"

	"github.com/bagheriali2001/GoWebDir/types"
)

func sortFiles(files []os.DirEntry) []os.DirEntry {
	dirList := []os.DirEntry{}
	fileList := []os.DirEntry{}

	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			dirList = append(dirList, files[i])
		} else {
			fileList = append(fileList, files[i])
		}
	}

	for i := 0; i < len(dirList); i++ {
		for j := i + 1; j < len(dirList); j++ {
			if dirList[i].Name() > dirList[j].Name() {
				dirList[i], dirList[j] = dirList[j], dirList[i]
			}
		}
	}

	for i := 0; i < len(fileList); i++ {
		for j := i + 1; j < len(fileList); j++ {
			if fileList[i].Name() > fileList[j].Name() {
				fileList[i], fileList[j] = fileList[j], fileList[i]
			}
		}
	}

	finalList := append(dirList, fileList...)

	return finalList
}

func ListDir(directory string) ([]types.File, error) {
	files, err := os.ReadDir(directory)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	var fileList []types.File

	files = sortFiles(files)

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}

		initSize := fileInfo.Size()

		var size float64
		var unit string
		switch {
		case initSize > 1<<30:
			size = float64(initSize) / (1 << 30)
			unit = "GB"
		case initSize > 1<<20:
			size = float64(initSize) / (1 << 20)
			unit = "MB"
		case initSize > 1<<10:
			size = float64(initSize) / (1 << 10)
			unit = "KB"
		default:
			size = float64(initSize)
			unit = "B"
		}
		size = float64(int(size*100)) / 100

		var fileType string

		if file.IsDir() {
			fileType = "Directory"
		} else {
			fileType = "File"
		}

		fileList = append(fileList, types.File{
			Name:    file.Name(),
			Size:    types.Size{Size: size, Unit: unit},
			Created: fileInfo.ModTime().Local().Format("2006-01-02 15:04:05"),
			Type:    fileType,
		})
	}

	return fileList, nil
}
