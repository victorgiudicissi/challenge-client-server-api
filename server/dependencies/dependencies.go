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
		log.Fatal(err)
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
