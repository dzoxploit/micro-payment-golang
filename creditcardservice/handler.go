package creditcardservice

import (
	"encoding/json"
	"net/http"
)

type CreditCardHandler struct {
	service CreditCardService
}

func NewCreditCardHandler(service CreditCardService) *CreditCardHandler {
	return &CreditCardHandler{
		service: service,
	}
}

func (h *CreditCardHandler) ValidateCreditCardHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		CreditCardNumber string `json:"credit_card_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	isValid := h.service.ValidateCreditCard(request.CreditCardNumber)

	response := struct {
		IsValid bool `json:"is_valid"`
	}{
		IsValid: isValid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
