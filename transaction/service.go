package transaction

import (
	"github.com/dzoxploit/micro-payment-golang/creditcardservice"
	"github.com/dzoxploit/micro-payment-golang/transactionhistory"
)

type Transaction struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type TransactionResponse struct {
	Status       string  `json:"status"`
	TotalPayment float64 `json:"total_payment"`
}

type TransactionService interface {
	ProcessTransaction(totalAmount float64, items []Transaction) TransactionResponse
}

type transactionService struct {
	creditCardService creditcardservice.CreditCardService
	historyService    transactionhistory.TransactionHistoryService
}

func NewTransactionService(creditCardService creditcardservice.CreditCardService, historyService transactionhistory.TransactionHistoryService) TransactionService {
	return &transactionService{
		creditCardService: creditCardService,
		historyService:    historyService,
	}
}

func (s *transactionService) ProcessTransaction(totalAmount float64, items []Transaction) TransactionResponse {
	// Perform credit card validation using creditCardService.ValidateCreditCard()
	// For simplicity, let's assume the credit card number is "4111 1111 1111 1111"
	isValid := s.creditCardService.ValidateCreditCard("4111111111111111")

	if !isValid {
		return TransactionResponse{
			Status: "failed",
		}
	}

	// Calculate the total payment by summing up the prices of all items
	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price
	}

	// Save the successful transaction to the history
	transaction := TransactionResponse{
		Status:       "success",
		TotalPayment: totalPrice,
	}
	s.historyService.SaveTransaction(transaction)

	response := TransactionResponse{
		Status:       "success",
		TotalPayment: totalPrice,
	}

	return response
}
