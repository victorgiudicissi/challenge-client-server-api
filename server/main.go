package main

import (
	"challeng-client-server-api/server/dependencies"
	quotationintegration "challeng-client-server-api/server/integration/quotation"
	quotationrepository "challeng-client-server-api/server/repository/quotation"
	"challeng-client-server-api/server/handler"
	"challeng-client-server-api/server/usecase"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	dep := dependencies.Init()

	defer dep.Destroy()

	quotationIntegration := quotationintegration.NewQuotationIntegration(dep.HttpClient)
	quotationRepository := quotationrepository.NewQuotationRepository(dep.DB)

	fetchQuotationUseCase := usecase.NewFetchQuotationUseCase(quotationRepository, quotationIntegration)

	fetchQuotationHandler := handler.NewQuotationHandler(fetchQuotationUseCase)

	r := mux.NewRouter()
	
	r.HandleFunc("/cotacao", fetchQuotationHandler.Handle)

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

	log.Println("Server started on port 8080")

	srv.ListenAndServe()
}
