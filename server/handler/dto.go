package handler

import "challeng-client-server-api/server/entities"

type QuotationResponse struct {
	Bid string `json:"bid"`
}

func NewQuotationResponse(e *entities.Quotation) *QuotationResponse {
	return &QuotationResponse{
		Bid: e.Bid,
	}
}