package dbsql

import (
	"database/sql"
	"reflect"
)

// Query  should query one col or panic
func (db *DB) QueryOne() (res []string, err error) {
	res = make([]string, 0)
	stmt, err := db.db.Prepare(db.SQLstring)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil || rows == nil {
		return
	}

	for rows.Next() {
		row := sql.NullString{}
		err = rows.Scan(&row)
		if err != nil {
			return
		}
		if row.Valid {
			res = append(res, row.String)
		} else {
			res = append(res, "")
		}

	}
	defer rows.Close()
	return
}

//Querys2Map only query like key value
func (db *DB) Querys2Map() (res map[string]string, err error) {
	res = make(map[string]string, 0)
	stmt, err := db.db.Prepare(db.SQLstring)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil || rows == nil {
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		return
	}
	for rows.Next() {
		arr := make([]interface{}, len(cols))
		// re := make([]string,len(cols))
		for i := 0; i < len(cols); i++ {
			arr[i] = new(sql.NullString)
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}
		if len(arr) == 2 {
			ktemp := reflect.ValueOf(arr[0])
			ktem := ktemp.Interface().(*sql.NullString)
			vtemp := reflect.ValueOf(arr[1])
			vtem := vtemp.Interface().(*sql.NullString)
			res[ktem.String] = vtem.String
		}
	}
	defer rows.Close()
	return
}

//Insert insert one row
func (db *DB) InsertOne(res []string) (err error) {
	conn, err := db.db.Begin()
	re := make([]interface{}, len(res))
	for i := 0; i < len(res); i++ {
		re[i] = res[i]
	}
	stmt, err := db.db.Prepare(db.SQLstring)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(re...)
	if err != nil {
		conn.Rollback()
		return
	}
	conn.Commit()
	defer stmt.Close()
	return
}
