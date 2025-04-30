package handler

import (
	"challeng-client-server-api/internal/usecase"
	"log"
)

type fetchQuotationHandler struct {
	uc usecase.FetchQuotationUseCase
}

func NewFetchQuotationHandler(uc usecase.FetchQuotationUseCase) *fetchQuotationHandler {
	return &fetchQuotationHandler{
		uc: uc,
	}
}

func (q *fetchQuotationHandler) Handle() error {
	err := q.uc.Do()

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
