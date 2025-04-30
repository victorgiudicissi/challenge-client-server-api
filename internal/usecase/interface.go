package usecase

import (
	"challeng-client-server-api/internal/entities"
)

type (
	QuotationRepository interface {
		SaveQuotation(quotation *entities.Quotation) error
	}

	QuotationIntegration interface {
		GetQuotation() (*entities.Quotation, error)
	}
)