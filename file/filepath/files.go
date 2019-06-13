package files

import (
	"os"
	"path"
	"path/filepath"
	"unicode"
)

type Files struct {
	ListFiles []string
}

//GetAllFile get files
func (files *Files) GetAllFile(pathsuffix ...string) (err error) {
	files.ListFiles = make([]string, 0)
	var (
		filespath string
		suffix    []string = make([]string, 0)
	)

	for k, v := range pathsuffix {
		if k == 0 {
			filespath = v
		} else {
			suffix = append(suffix, v)
		}
	}

	err = filepath.Walk(filespath, func(filespath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		
		fileSuffix := path.Ext(filespath)

		if len(suffix) == 0 {
			files.ListFiles = append(files.ListFiles, filespath)
		} else {
			for _, v := range suffix {
				if fileSuffix == v {
					files.ListFiles = append(files.ListFiles, filespath)
				}
			}
		}

		return err
	})

	return
}

func (files *Files) FilefilterisChineseChar(filter string) {
	fp := files.ListFiles
	files.ListFiles = make([]string, 0)

	for _, p := range fp {
		if !isChineseChar(p, filter) {
			files.ListFiles = append(files.ListFiles, p)
		}
		// for _, r := range p {
		// 	if unicode.Is(unicode.Scripts[filter], r) {
		// 		files.ListFiles = append(files.ListFiles, p)
		// 	}
		// }
	}
}

func isChineseChar(str string, filter string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts[filter], r) {
			return true
		}
	}
	return false
}
