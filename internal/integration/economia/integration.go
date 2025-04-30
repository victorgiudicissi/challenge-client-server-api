package economia

import (
	"net/http"

	"challeng-client-server-api/internal/usecase"
)

type economiaIntegration struct {
	client *http.Client
}

func NewEconomiaIntegration(client *http.Client) usecase.EconomiaIntegration {
	return &economiaIntegration{
		client: client,
	}
}
