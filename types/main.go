package types

type Size struct {
	Size float64
	Unit string
}

type File struct {
	Name    string
	Size    Size
	Created string
	Type    string
}

type Directory struct {
	Name  string
	Value string
}

type PageData struct {
	Directory Directory
	Files     []File
	IsRoot    bool
}

type HandlerWrapper struct {
	RootPath          string
	ShowHiddenFiles   bool
	ShowHiddenFolders bool
}
