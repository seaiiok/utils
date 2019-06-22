package fmtx

import (
	"fmt"
	"runtime"
	"syscall"
)

const (
	//蓝色
	Ok = 2 * (iota + 1)
	//红色
	Err
	//黄色
	Warn
	//灰色
	Info
)

func Println(i int, a ...interface{}) (int, error) {
	if IsWindows() {
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

func Printf(i int, format string, a ...interface{}) (int, error) {
	if IsWindows() {
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

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
