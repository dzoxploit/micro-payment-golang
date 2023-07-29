package transaction

import (
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service TransactionService
}

func NewTransactionHandler(service TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) ProcessTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		TotalAmount float64       `json:"total_amount"`
		Items       []Transaction `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	response := h.service.ProcessTransaction(request.TotalAmount, request.Items)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
