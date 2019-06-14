package util

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"
)

type File struct {
	files
}

type files struct {
	FilesList []string
	Err       error
}

func (fi *File) GetAllFiles(pathstring string) *files {
	fi.files.FilesList = make([]string, 0)

	err := filepath.Walk(pathstring, func(pathstring string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		fi.files.FilesList = append(fi.files.FilesList, pathstring)
		return nil
	})
	fi.files.Err = err
	return &fi.files
}

func (fi *files) GetFilterSuffix(suffix ...string) *files {
	if len(suffix) == 0 {
		return fi
	} else {
		for fk, fv := range fi.FilesList {
			fileSuffix := path.Ext(fv)

			for _, sv := range suffix {
				if strings.EqualFold(fileSuffix, sv) {
					fi.FilesList=append(fi.FilesList[:fk],fi.FilesList[fk+1:]...)
				}
			}
		}
	}
	return fi
}

func (fi *files) GetFilterZH() *files {
	for k, v := range fi.FilesList {
		if isChineseChar(v) {
			fi.FilesList=append(fi.FilesList[:k],fi.FilesList[k+1:]...)
			
		}
	}
	return fi
}

func isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
