package handler

import "challeng-client-server-api/internal/entities"

type QuotationResponse struct {
	Bid string `json:"bid"`
}

func NewQuotationResponse(e *entities.Quotation) *QuotationResponse {
	return &QuotationResponse{
		Bid: e.Bid,
	}
}