package dependencies

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Dependencies struct {
	DB         *sql.DB
	HttpClient *http.Client
}

func Init() *Dependencies {

	return &Dependencies{
		DB:         initDB(),
		HttpClient: initHttpClient(),
	}
}

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

func initHttpClient() *http.Client {
	return &http.Client{Timeout: 200 * time.Millisecond}
}

func (d *Dependencies) Destroy() {
	if err := d.DB.Close(); err != nil {
		log.Fatal(err)
	}
}
