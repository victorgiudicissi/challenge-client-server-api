package quotation

import (
	"challeng-client-server-api/server/usecase"
	"database/sql"
)

type quotationRepository struct {
}

func NewQuotationRepository(_ *sql.DB) usecase.QuotationRepository {
	return &quotationRepository{}
}
