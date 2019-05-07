package dbsql

import (
	"database/sql"
	"sync"
)

//DB a database struct
type DB struct {
	DriverName     string
	DataSourceName string
	SQLstring      string
	db             *sql.DB
	once           sync.Once
}
