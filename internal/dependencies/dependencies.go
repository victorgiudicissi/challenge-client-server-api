package dependencies

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type ServerDependencies struct {
	DB         *sql.DB
	HttpClient *http.Client
}

type ClientDependencies struct {
	HttpClient *http.Client
}

func InitServer() *ServerDependencies {
	return &ServerDependencies{
		DB:         initDB(),
		HttpClient: initHttpClient(),
	}
}

func InitClient() *ClientDependencies {
	return &ClientDependencies{
		HttpClient: initHttpClient(),
	}
}

func (d *ServerDependencies) Close() {
	if err := d.DB.Close(); err != nil {
		log.Fatal(err)
	}
}
