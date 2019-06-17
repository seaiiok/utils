package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {

	DriverName := "mysql"
	DataSourceName := "root:1949@tcp(127.0.0.1:3306)/ifixsvr"
	
		db,err:=New(DriverName, DataSourceName)
		if err != nil {
			t.Log(err)
		}
	

}
