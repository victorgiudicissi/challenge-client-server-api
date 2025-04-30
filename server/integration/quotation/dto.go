package quotation

import "challeng-client-server-api/server/entities"

type QuotationResponse struct {
	USDBRL Coin `json:"USDBRL"`
}

type Coin struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func (q QuotationResponse) ToEntity() *entities.Quotation {
	return &entities.Quotation{
		Bid: q.USDBRL.Bid,
	}
}
