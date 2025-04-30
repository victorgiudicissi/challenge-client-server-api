package dependencies

import "database/sql"

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./quotation.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS quotation (
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			code 		TEXT NOT NULL,
			codein 		TEXT NOT NULL,
			bid			TEXT NOT NULL,
			created_at 	TEXT NOT NULL
		)`,
	)

	if err != nil {
		panic(err)
	}

	return db
}
