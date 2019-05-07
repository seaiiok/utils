package dbsql

import (
	"database/sql"
	"reflect"
)

//Init mssql driver init
func (db *DB) Init() (err error) {
	db.once.Do(func() {
		db.db, err = sql.Open(db.DriverName, db.DataSourceName)
	})
	return db.db.Ping()
}

// Create create table
func (db *DB) Create(databaseName string) (err error) {
	return nil
}

// Query  query more
func (db *DB) Query() (res [][]string, err error) {
	res = make([][]string, 0)
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

//Insert insert more rows
func (db *DB) Insert(res [][]string) (err error) {
	conn, err := db.db.Begin()

	for c := 0; c < len(res); c++ {
		re := make([]interface{}, len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, err := db.db.Prepare(db.SQLstring)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
		_, err = stmt.Exec(re...)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
	}
	conn.Commit()
	return
}

//Update update more rows
func (db *DB) Update(res [][]string) (err error) {
	conn, err := db.db.Begin()

	for c := 0; c < len(res); c++ {
		re := make([]interface{}, len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, err := db.db.Prepare(db.SQLstring)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
		_, err = stmt.Exec(re...)
		if err != nil {
			conn.Rollback()
		}
		defer stmt.Close()
	}
	conn.Commit()
	return
}

//Delete delete more rows
func (db *DB) Delete() (err error) {
	conn, err := db.db.Begin()
	stmt, err := db.db.Prepare(db.SQLstring)
	if err != nil {
		conn.Rollback()
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		conn.Rollback()
	}
	defer stmt.Close()
	conn.Commit()
	return
}
