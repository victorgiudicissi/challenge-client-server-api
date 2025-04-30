package file

import (
	"challeng-client-server-api/internal/usecase"
)

type quotationRepository struct {
}

func NewQuotationRepository() usecase.QuotationRepository {
	return &quotationRepository{
	}
}
