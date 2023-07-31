package main

import (
	"log"
	"net/http"

	"github.com/dzoxploit/micro-payment-golang/creditcardservice" // Corrected import path
	"github.com/dzoxploit/micro-payment-golang/transactions"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize services
	creditCardService := creditcardservice.NewCreditCardService()
	transactionHistoryService := transactions.NewTransactionHistoryService()

	// Initialize handlers
	creditCardHandler := creditcardservice.NewCreditCardHandler(creditCardService)
	transactionHandler := transactions.NewTransactionHandler(transactions.NewTransactionService(creditCardService, transactionHistoryService))
	transactionHistoryHandler := transactions.NewTransactionHistoryHandler(transactionHistoryService)

	// Create a new router
	router := mux.NewRouter()

	// Credit Card Service endpoints
	router.HandleFunc("/creditcard/validate", creditCardHandler.ValidateCreditCardHandler).Methods("POST")

	// Transaction Service endpoints
	router.HandleFunc("/transaction/process", transactionHandler.ProcessTransactionHandler).Methods("POST")

	// Transaction History Service endpoints
	router.HandleFunc("/transactionhistory/save", transactionHistoryHandler.SaveTransactionHandler).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":7000", router))
}
