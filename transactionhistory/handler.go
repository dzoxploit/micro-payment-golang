package transactionhistory

import (
	"encoding/json"
	"net/http"

	"github.com/dzoxploit/micro-payment-golang/transaction"
)

type TransactionHistoryHandler struct {
	service TransactionHistoryService
}

func NewTransactionHistoryHandler(service TransactionHistoryService) *TransactionHistoryHandler {
	return &TransactionHistoryHandler{
		service: service,
	}
}

func (h *TransactionHistoryHandler) SaveTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Transaction transaction.Transaction `json:"transaction"` // Use transaction.Transaction from the imported package
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	h.service.SaveTransaction(request.Transaction)

	w.WriteHeader(http.StatusOK)
}
