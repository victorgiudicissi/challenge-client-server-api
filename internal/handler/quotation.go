package handler

import (
	"challeng-client-server-api/internal/usecase"
	"encoding/json"
	"net/http"
)

type quotationHandler struct {
	uc usecase.FetchQuotationUseCase
}

func NewQuotationHandler(uc usecase.FetchQuotationUseCase) *quotationHandler {
	return &quotationHandler{
		uc: uc,
	}
}

func (q *quotationHandler) Handle(w http.ResponseWriter, r *http.Request) {
	result, err := q.uc.Do()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(NewQuotationResponse(result))
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
