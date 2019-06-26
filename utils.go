package utils

type utils struct {
	Cmd   _cmd
	Os    _os
	Times _times
}

//New ...
//
//控制台执行类 Cmd
//
//系统操作类 Os
//
//时间操作类 Times
//
func New() *utils {
	return &utils{}
}
