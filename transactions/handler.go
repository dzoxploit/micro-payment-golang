package transactions

import (
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service TransactionsService
}

type TransactionHistoryHandler struct {
	service TransactionHistoryService
}


func NewTransactionHandler(service TransactionsService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) ProcessTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		TotalAmount float64       `json:"total_amount"`
		Items       []Transactions `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	response := h.service.ProcessTransaction(request.TotalAmount, request.Items)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func NewTransactionHistoryHandler(service TransactionHistoryService) *TransactionHistoryHandler {
	return &TransactionHistoryHandler{
		service: service,
	}
}

func (h *TransactionHistoryHandler) SaveTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Transactions Transactions `json:"transaction"` // Use transaction.Transaction from the imported package
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.service.SaveTransaction(request.Transactions)

	w.WriteHeader(http.StatusOK)
}
