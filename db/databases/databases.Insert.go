package databases

//Insert insert more rows
func (db *DB) Insert(SQLstring string, res [][]string) (err error) {
	conn, err := db.db.Begin()

	for c := 0; c < len(res); c++ {
		re := make([]interface{}, len(res[c]))
		for i := 0; i < len(res[c]); i++ {
			re[i] = res[c][i]
		}
		stmt, err := db.db.Prepare(SQLstring)
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
