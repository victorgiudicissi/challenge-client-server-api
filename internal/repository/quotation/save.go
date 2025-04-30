package quotation

import (
	"challeng-client-server-api/internal/entities"
	"log"
	"time"
)

func (q *quotationRepository) SaveQuotation(quotation *entities.Quotation) error {
	now := time.Now().Format(time.RFC3339)

	stmt, err := q.db.Prepare(`INSERT INTO quotation (code, codein, bid, created_at) VALUES (?, ?, ?, ?)`)
	
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close() 

	_, err = stmt.Exec(quotation.Code, quotation.Codein, quotation.Bid, now)
	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}
