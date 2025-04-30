package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	// URL da API de câmbio
	APIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/cotacao", CotacaoHandler)

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
	log.Println("Server started on port 8080")

	srv.ListenAndServe()
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	c := http.Client{Timeout:  200 * time.Millisecond}

	resp, err := c.Get(APIURL)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler o corpo da resposta:", err)
		return
	}

	var data CambioResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	log.Println("Dados recebidos:", data)

	w.WriteHeader(http.StatusOK)
}
