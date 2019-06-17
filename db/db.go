package db

import (
	"database/sql"
	"sync"
)

var once sync.Once

//DB a database struct
type DB struct {
	db *sql.DB
}

//New return DB
func New(DriverName, DataSourceName string) (db *DB, err error) {
	db = new(DB)
	once.Do(func() {
		db.db, err = sql.Open(DriverName, DataSourceName)
		err = db.db.Ping()
	})
	return
}




