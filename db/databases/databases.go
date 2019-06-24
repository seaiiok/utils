package databases

import (
	"database/sql"
)

//NewDB return DB
func NewDB(DriverName, DataSourceName string) (db *DB, err error) {
	db = new(DB)
	db.once.Do(func() {
		db.db, err = sql.Open(DriverName, DataSourceName)
	})
	err = db.db.Ping()
	return
}
