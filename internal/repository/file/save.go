package file

import (
	"challeng-client-server-api/internal/entities"
	"fmt"
	"log"
	"os"
)

func (q *quotationRepository) SaveQuotation(quotation *entities.Quotation) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data := fmt.Sprintf("Dólar: %s", quotation.Bid)

	_, err = file.WriteString(data)
	if err != nil {
		log.Println(err)
	}
	
	return nil
}
