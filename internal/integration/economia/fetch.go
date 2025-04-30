package economia

import (
	"challeng-client-server-api/internal/entities"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	APIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func (i *economiaIntegration) FetchQuotation() (*entities.Quotation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
    defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, APIURL, nil)
    if err != nil {
        return nil, err
	}

	resp, err := i.client.Do(req)
    if err != nil {
        return nil, err
    }

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result *QuotationResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.ToEntity(), nil
}
