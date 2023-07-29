package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/dzoxploit/micro-payment-golang/creditcardservice"
	"github.com/dzoxploit/micro-payment-golang/transaction"
	"github.com/dzoxploit/micro-payment-golang/transactionhistory" // Corrected import path
)

func main() {
	// Initialize services
	creditCardService := creditcardservice.NewCreditCardService()
	transactionHistoryService := transactionhistory.NewTransactionHistoryService()

	// Initialize handlers
	creditCardHandler := creditcardservice.NewCreditCardHandler(creditCardService)
	transactionHandler := transaction.NewTransactionHandler(transaction.NewTransactionService(creditCardService, transactionHistoryService))
	transactionHistoryHandler := transactionhistory.NewTransactionHistoryHandler(transactionHistoryService)

	// Create a new router
	router := mux.NewRouter()

	// Credit Card Service endpoints
	router.HandleFunc("/creditcard/validate", creditCardHandler.ValidateCreditCardHandler).Methods("POST")

	// Transaction Service endpoints
	router.HandleFunc("/transaction/process", transactionHandler.ProcessTransactionHandler).Methods("POST")

	// Transaction History Service endpoints
	router.HandleFunc("/transactionhistory/save", transactionHistoryHandler.SaveTransactionHandler).Methods("POST")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
