package quotation

import (
	"challeng-client-server-api/internal/entities"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	APIURL = "http://localhost:8080/cotacao"
)

func (i *quotationIntegration) FetchQuotation() (*entities.Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, APIURL, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var result *QuotationResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
	}

	return result.ToEntity(), nil
}
