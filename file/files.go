package files

import (
	"utils/file/filepath"
)

//NewGetFiles get files
func NewGetFiles() *files.Files {
	return &files.Files{
		ListFiles: make([]string, 0),
	}
}
