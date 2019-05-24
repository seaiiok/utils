package databases

import (
	"database/sql"
	"reflect"
)

// Query  query more
func (db *DB) QueryStr(SQLstring string) (res [][]string, err error) {
	res = make([][]string, 0)
	stmt, err := db.db.Prepare(SQLstring)
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
		re := make([]string, len(cols))
		for i := 0; i < len(cols); i++ {
			arr[i] = new(sql.NullString)
		}
		err = rows.Scan(arr...)
		if err != nil {
			return
		}

		for i := 0; i < len(arr); i++ {
			arrtemp := reflect.ValueOf(arr[i])
			arrtem := arrtemp.Interface().(*sql.NullString)
			re[i] = arrtem.String
		}
		res = append(res, re)
	}
	defer rows.Close()
	return
}

// Query  query more
func (db *DB) Query(SQLstring string) (res [][]interface{}, err error) {
	res = make([][]interface{}, 0)
	stmt, err := db.db.Prepare(SQLstring)
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
		err = rows.Scan(arr...)
		if err != nil {
			return
		}

		res = append(res, arr)
	}
	defer rows.Close()
	return
}
