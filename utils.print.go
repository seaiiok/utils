package utils

import (
	"fmt"
	"syscall"
)

/**
* 常用颜色
* 蓝色 1
* 红色 3
* 黄色 5
* 灰色 7
**/
type print struct {
	_os os
}

//终端输出,增加颜色参数，类似fmt.Println
func (p *print) Println(color int, a ...interface{}) (int, error) {
	if p._os.IsWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
		n, err := fmt.Println(a...)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return n, err
	} else {
		return fmt.Println(a...)
	}

}

//终端格式输出,增加颜色参数，类似fmt.Printf
func (p *print) Printf(color int, format string, a ...interface{}) (int, error) {
	if p._os.IsWindows() {
		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		proc := kernel32.NewProc("SetConsoleTextAttribute")
		handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(color))
		n, err := fmt.Printf(format, a...)
		handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
		CloseHandle := kernel32.NewProc("CloseHandle")
		CloseHandle.Call(handle)
		return n, err
	} else {
		return fmt.Println(fmt.Sprintf(format, a...))
	}

}
