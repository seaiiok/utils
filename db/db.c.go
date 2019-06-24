package db

// Create create table
func (db *DB) Create(SQLstring string) (err error) {
	conn, err := db.db.Begin()
	stmt, err := db.db.Prepare(SQLstring)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		conn.Rollback()
		return
	}
	conn.Commit()
	defer stmt.Close()
	return
}
