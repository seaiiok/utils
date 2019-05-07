package db

//DB interface
type DB interface {
	Init() error
	Create(string) error
	Query() [][]string
	Insert([]string) error
	Update([]string) error
	Delete() error
}
