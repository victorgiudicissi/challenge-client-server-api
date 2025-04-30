package quotation

import "challeng-client-server-api/internal/entities"

type QuotationResponse struct {
	Code        string `json:"code"`
	Codein      string `json:"codein"`
	Bid         string `json:"bid"`
	CreatedDate string `json:"created_date"`
}

func (q *QuotationResponse) ToEntity() *entities.Quotation {
	return &entities.Quotation{
		Bid:    q.Bid,
		Code:   q.Code,
		Codein: q.Codein,
	}
}
