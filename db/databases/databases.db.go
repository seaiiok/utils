package databases

import (
	"database/sql"
	"sync"
)

//DB a database struct
type DB struct {
	db   *sql.DB
	once sync.Once
}
