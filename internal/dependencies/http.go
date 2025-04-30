package dependencies

import (
	"net/http"
)

func initHttpClient() *http.Client {
	return &http.Client{}
}
