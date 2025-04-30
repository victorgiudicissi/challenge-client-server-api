package main

import (
	"challeng-client-server-api/internal/dependencies"
	"challeng-client-server-api/internal/handler"
	quotationintegration "challeng-client-server-api/internal/integration/quotation"
	quotationsqliterepository "challeng-client-server-api/internal/repository/file"
	"challeng-client-server-api/internal/usecase"
	"log"
)

func main() {
	dep := dependencies.InitClient()

	quotationIntegration := quotationintegration.NewQuotationIntegration(dep.HttpClient)
	quotationRepository := quotationsqliterepository.NewQuotationRepository()

	fetchQuotationUseCase := usecase.NewFetchQuotationUseCase(quotationRepository, quotationIntegration)

	fetchQuotationHandler := handler.NewFetchQuotationHandler(fetchQuotationUseCase)

	err := fetchQuotationHandler.Handle()
	
	if err != nil {
		log.Fatalf("Error handling fetch quotation: %v", err)
	}

	log.Println("Data fetched and saved successfully.")
}
