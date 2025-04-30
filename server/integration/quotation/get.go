package quotation

import (
	"challeng-client-server-api/server/entities"
	"encoding/json"
	"io"
)

const (
	APIURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func (i *quotationIntegration) GetQuotation() (*entities.Quotation, error) {
	resp, err := i.client.Get(APIURL)
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
