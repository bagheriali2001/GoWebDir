package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorReset = "\033[0m"

func PrintStartupConfig(port int, folderPath string, showHiddenFiles, showHiddenFolders bool) {
	// print the configuration with a nice format and colors
	fmt.Fprintf(os.Stdout, "%sConfiguration:%s\n", colorYellow, colorReset)
	fmt.Fprintf(os.Stdout, "\t%sPort:%s %d\n", colorGreen, colorReset, port)
	fmt.Fprintf(os.Stdout, "\t%sServing files from:%s %s\n", colorGreen, colorReset, folderPath)
	fmt.Fprintf(os.Stdout, "\t%sShow Hidden Files:%s %t\n", colorGreen, colorReset, showHiddenFiles)
	fmt.Fprintf(os.Stdout, "\t%sShow Hidden Folders:%s %t\n\n\n", colorGreen, colorReset, showHiddenFolders)
}

func HTTPLog(r *http.Request, status int) {
	fmt.Fprintf(os.Stdout, "%s%s %s%s %s%s %s%d\n", colorYellow, time.Now().Format("2006-01-02 15:04:05"), colorRed, r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")], colorReset, r.URL.Path, colorGreen, status)
}
