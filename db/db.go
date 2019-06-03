package db

import (
	"github.com/seaiiok/utils/db/databases"
)

//NewDB new database driver
func NewDB(DriverName, DataSourceName string) (*databases.DB, error) {
	return databases.NewDB(DriverName, DataSourceName)
}
