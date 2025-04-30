package sqlite

import (
	"challeng-client-server-api/internal/entities"
	"context"
	"log"
	"time"
)

func (q *quotationRepository) SaveQuotation(quotation *entities.Quotation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
    defer cancel()
	
	now := time.Now().Format(time.RFC3339)

	stmt, err := q.db.Prepare(`INSERT INTO quotation (code, codein, bid, created_at) VALUES (?, ?, ?, ?)`)
	
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close() 

	_, err = stmt.ExecContext(ctx, quotation.Code, quotation.Codein, quotation.Bid, now)
	if err != nil {
		log.Println(err)
	}
	
	return nil
}
