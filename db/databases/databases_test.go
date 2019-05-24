package databases_test

import (
	"testing"
	"utils/db/databases"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewDB(t *testing.T) {
	DriverName := "mysql"
	DataSourceName := "root:1949@/test"
	db, err := databases.NewDB(DriverName, DataSourceName)
	if err != nil {
		t.Log(err)
	}
	t.Log(db)
}

// func TestCreate(t *testing.T) {
// 	var d *db.DB
// 	d = new(db.DB)
// 	d.DriverName = "sqlite3"
// 	d.DataSourceName = "D:\\sqlite3\\DB\\mygit.db"
// 	err := d.Init()
// 	t.Log(err)

// 	d.SQLstring = `CREATE TABLE mygit ( ID TEXT PRIMARY KEY NOT NULL, NAME TEXT );`
// 	err = d.Create()
// 	t.Log(err)
// }

func TestInsert(t *testing.T) {
	DriverName := "mysql"
	DataSourceName := "root:1949@/test"
	db, err := databases.NewDB(DriverName, DataSourceName)
	if err != nil {
		t.Log(err)
	}
	t.Log(db)

	SQLstring := `insert into T1 (id,name) values(?,?);`
	res := make([][]interface{}, 0, 0)
	re := make([]interface{}, 2)
	re[0] = "1"
	re[1] = "Tome"
	res = append(res, re)
	re1 := make([]interface{}, 2)
	re1[0] = "2"
	re1[1] = "Jime"
	res = append(res, re1)
	t.Log(res)
	err = db.Insert2(SQLstring, res)
	t.Log(err)
}

// func TestUpdate(t *testing.T) {
// 	var d *db.DB
// 	d = new(db.DB)
// 	d.DriverName = "sqlite3"
// 	d.DataSourceName = "D:\\sqlite3\\DB\\mygit.db"
// 	err := d.Init()
// 	t.Log(err)

// 	d.SQLstring = `update mygit set NAME=? where ID=?`
// 	res := make([][]string, 0, 0)
// 	re := make([]string, 2)
// 	re[0] = "Tom"
// 	re[1] = "1"
// 	res = append(res, re)
// 	re1 := make([]string, 2)
// 	re1[0] = "Jim"
// 	re1[1] = "2"
// 	res = append(res, re1)
// 	t.Log(res)
// 	err = d.Update(res)
// 	t.Log(err)
// }

// func TestQuery(t *testing.T) {
// 	var d *db.DB
// 	d = new(db.DB)
// 	d.DriverName = "sqlite3"
// 	d.DataSourceName = "D:\\sqlite3\\DB\\mygit.db"
// 	err := d.Init()
// 	t.Log(err)

// 	d.SQLstring = `select * from mygit where ID="1";`

// 	res, err := d.Query()
// 	fmt.Println("查询:",res)
// 	t.Log(err)
// }

// func BenchmarkQuery(t *testing.B) {
// 	var d *db.DB
// 	d = new(db.DB)
// 	d.DriverName = "sqlite3"
// 	d.DataSourceName = "D:\\sqlite3\\DB\\mygit.db"
// 	err := d.Init()
// 	t.Log(err)

// 	d.SQLstring = `select * from mygit where ID="1";`

// 	for i := 0; i < t.N; i++ {
// 		d.Query()
// 	}
// 	t.Log(err)
// }

// func TestDelete(t *testing.T) {
// 	var d *db.DB
// 	d = new(db.DB)
// 	d.DriverName = "sqlite3"
// 	d.DataSourceName = "D:\\sqlite3\\DB\\mygit.db"
// 	err := d.Init()
// 	t.Log(err)
// // DELETE FROM COMPANY WHERE ID = 7;
// 	d.SQLstring = `DELETE FROM mygit WHERE ID = ?;`
// 	res := make([][]string, 0, 0)
// 	re := make([]string, 1)
// 	re[0] = "1"
// 	res = append(res, re)
// 	// re1 := make([]string, 1)
// 	// re1[0] = "2"
// 	// res = append(res, re1)
// 	t.Log(res)
// 	err = d.Delete(res)
// 	t.Log(err)
// }
