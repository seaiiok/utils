package utils

import (
	"runtime"
)

type os struct {
}

//判断目标是否为微软Windows系统
func (o *os) IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

//查看目标操作系统
func (o *os) OsName() string {
	return runtime.GOOS
}
