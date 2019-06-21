package db

import (
	"database/sql"
	"fmt"
	"strconv"
)

//Result...
type Result struct {
	Cols []string
	Res  [][]interface{}
	Err  error
}

type col struct {
	colName string
}

// Query ...
func (db *DB) Query(SQLstring string) (res Result) {
	res.Res = make([][]interface{}, 0)

	stmt, err := db.db.Prepare(SQLstring)
	if err != nil {
		res.Err = err
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil || rows == nil {
		res.Err = err
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		res.Err = err
		return
	}

	colstype, err := rows.ColumnTypes()
	if err != nil {
		res.Err = err
		return
	}
	for _, v := range colstype {
		fmt.Println(v.DatabaseTypeName())
		fmt.Println(v.DecimalSize())
		fmt.Println(v.Length())
		fmt.Println(v.Name())
		fmt.Println(v.Nullable())
		fmt.Println(v.ScanType())
	}

	res.Cols = make([]string, 0)
	res.Cols = append(res.Cols, cols...)

	values := make([]sql.RawBytes, len(cols))
	scans := make([]interface{}, len(cols))

	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		row := make([]interface{}, len(cols))
		err = rows.Scan(scans...)
		if err != nil {
			return
		}
		for k, v := range values {
			if colstype[k].DatabaseTypeName() == "INT" {
				row[k], _ = strconv.Atoi(string(v))
			} else if colstype[k].DatabaseTypeName() == "VARCHAR" {
				row[k] = string(v)
			}
			// row[k] = string(v)
		}
		res.Res = append(res.Res, row)
	}
	defer rows.Close()
	return
}

// Query ...
// func (db *DB) query(SQLstring string) (rows *sql.Rows, err error) {
// 	stmt, err := db.db.Prepare(SQLstring)
// 	if err != nil {
// 		res.Err = err
// 		return
// 	}
// 	defer stmt.Close()
// 	rows, err = stmt.Query()
// 	if err != nil || rows == nil {
// 		res.Err = err
// 		return
// 	}
// 	defer rows.Close()
// 	return
// }

// func (res Result) ConvToArray() (resArr [][]string) {
// 	resArr = make([][]string, len(res.Res))

// 	for _, resv := range res.Res {
// 		reArr := make([]string, len(resv))
// 		for _, rev := range resv {
// 			x := interface2String(rev)
// 			reArr = append(reArr, x)
// 		}
// 		resArr = append(resArr, reArr)
// 	}

// 	return
// }

func (res Result) ConvToMap() []map[string]string {
	re := make([]map[string]string, 0)
	return re
}
