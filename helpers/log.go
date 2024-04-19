package helpers

import (
	"fmt"
	"os"
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
