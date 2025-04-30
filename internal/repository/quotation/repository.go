package quotation

import (
	"challeng-client-server-api/internal/usecase"
	"database/sql"
)

type quotationRepository struct {
	db *sql.DB
}

func NewQuotationRepository(db *sql.DB) usecase.QuotationRepository {
	return &quotationRepository{
		db: db,
	}
}
