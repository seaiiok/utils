package db

import (
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {

	DriverName := "mysql"
	DataSourceName := "root:1949@tcp(127.0.0.1:3306)/test"

	db, err := New(DriverName, DataSourceName)
	if err != nil {
		t.Log(err)
	}
	res := db.Query("select id as iid ,age as aage from t2")
	t.Log(res)

	for _, v := range res.Res {
		for _, rv := range v {
			t.Log(rv, reflect.TypeOf(rv))
		}

	}

}
