package quotation

import (
	"net/http"

	"challeng-client-server-api/server/usecase"
)

type quotationIntegration struct {
	client *http.Client
}

func NewQuotationIntegration(client *http.Client) usecase.QuotationIntegration {
	return &quotationIntegration{
		client: client,
	}
}
