package quotation

import (
	"net/http"

	"challeng-client-server-api/internal/usecase"
)

type quotationIntegration struct {
	client *http.Client
}

func NewQuotationIntegration(client *http.Client) usecase.QuotationIntegration {
	return &quotationIntegration{
		client: client,
	}
}
