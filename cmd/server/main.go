package main

import (
	"challeng-client-server-api/internal/dependencies"
	"challeng-client-server-api/internal/handler"
	economiaintegration "challeng-client-server-api/internal/integration/economia"
	quotationsqliterepository "challeng-client-server-api/internal/repository/sqlite"
	"challeng-client-server-api/internal/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	dep := dependencies.InitServer()

	defer dep.Close()

	quotationIntegration := economiaintegration.NewEconomiaIntegration(dep.HttpClient)
	quotationRepository := quotationsqliterepository.NewQuotationRepository(dep.DB)

	fetchQuotationUseCase := usecase.NewFetchAndSaveQuotationUseCase(quotationRepository, quotationIntegration)

	fetchQuotationHandler := handler.NewFetchAndSaveQuotationHandler(fetchQuotationUseCase)

	r := mux.NewRouter()

	r.HandleFunc("/cotacao", fetchQuotationHandler.Handle).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port 8080")

	srv.ListenAndServe()
}
