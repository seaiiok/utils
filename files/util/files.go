package util

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
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
	i := 0
	if len(suffix) == 0 {
		return fi
	} else {
		for fk, fv := range fi.FilesList {
			fileSuffix := path.Ext(fv)

			for _, sv := range suffix {
				if strings.EqualFold(fileSuffix, sv) {
					fi.FilesList[i] = fi.FilesList[fk]
					i++
					continue
				}
			}
		}
	}
	fi.FilesList = fi.FilesList[:i]
	return fi
}

func (fi *files) GetFilterReg(r ...string) *files {
	i := 0
	for k, v := range fi.FilesList {
		for _, rv := range r {
			reg := regexp.MustCompile(rv)
			if reg.MatchString(v) {
				fi.FilesList[i] = fi.FilesList[k]
				i++
				continue
			}
		}
	}
	fi.FilesList = fi.FilesList[:i]
	return fi
}

func (fi *files) GetFilterZH() *files {
	i := 0
	for k, v := range fi.FilesList {
		if !isChineseChar(v) {
			fi.FilesList[i] = fi.FilesList[k]
			i++
			continue
		}
	}
	fi.FilesList = fi.FilesList[:i]
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
