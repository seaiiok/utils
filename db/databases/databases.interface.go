package databases

//dbInterface DB interface
type dbInterface interface {
	//创建数据库表
	Create() error

	//查询数据，返回二维数据结果集
	Query() [][]string

	//插入数据，数据格式要求为数组
	Insert([]string) error

	//更新数据，数据格式要求为数组
	Update([]string) error

	//删除数据
	Delete() error
}
