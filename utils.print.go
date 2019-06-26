package utils

import (
	"fmt"
	"runtime"
	"syscall"
)

/**
* 常用颜色
* 蓝色 1
* 红色 3
* 黄色 5
* 灰色 7
**/
type print struct{}

func (p *print) Println(i int, a ...interface{}) (int, error) {
	if p.IsWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
		n, err := fmt.Println(a...)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return n, err
	} else {
		return fmt.Println(a...)
	}

}

func (p *print) Printf(i int, format string, a ...interface{}) (int, error) {
	if p.IsWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
		n, err := fmt.Println(fmt.Sprintf(format, a...))
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return n, err
	} else {
		return fmt.Println(fmt.Sprintf(format, a...))
	}

}

func (p *print) IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
