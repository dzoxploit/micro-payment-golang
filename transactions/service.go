package transactions

import (
	"log"

	"github.com/dzoxploit/micro-payment-golang/creditcardservice"
	"github.com/dzoxploit/micro-payment-golang/database"
)

type Transactions struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type TransactionsService interface {
	ProcessTransaction(totalAmount float64, items []Transactions) TransactionsResponse
}

type TransactionsResponse struct {
	Status       string  `json:"status"`
	TotalPayment float64 `json:"total_payment"`
}

type transactionsService struct {
	creditCardService creditcardservice.CreditCardService
	historyService    TransactionHistoryService
}
type TransactionHistoryService interface {
	SaveTransaction(transaction Transactions)
}

type transactionHistoryService struct{}

func NewTransactionHistoryService() TransactionHistoryService {
	return &transactionHistoryService{}
}

func (s *transactionHistoryService) SaveTransaction(transaction Transactions) {
	// Buka koneksi ke database MySQL
	err := database.InitDB()
	if err != nil {
		log.Printf("Failed to connect to the database: %v\n", err)
		return
	}
	defer database.CloseDB()

	// SQL statement untuk menyimpan data transaksi ke dalam tabel "transactions"
	stmt := "INSERT INTO transactions (name, price) VALUES (?, ?)"

	// Eksekusi SQL statement dengan data transaksi
	_, err = database.DBConn.Exec(stmt, transaction.Name, transaction.Price)
	if err != nil {
		log.Printf("Failed to save transaction to the database: %v\n", err)
		return
	}

	log.Println("Transaction successfully saved to the database.")
}


func NewTransactionService(creditCardService creditcardservice.CreditCardService, historyService TransactionHistoryService) TransactionsService {
	return &transactionsService{
		creditCardService: creditCardService,
		historyService:    historyService,
	}
}

func (s *transactionsService) ProcessTransaction(totalAmount float64, items []Transactions) TransactionsResponse {
	// Perform credit card validation using creditCardService.ValidateCreditCard()
	// For simplicity, let's assume the credit card number is "4111 1111 1111 1111"
	isValid := s.creditCardService.ValidateCreditCard("4111111111111111")

	if !isValid {
		return TransactionsResponse{
			Status: "failed",
		}
	}

	// Calculate the total payment by summing up the prices of all items
	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price
	}

	// Save the successful transaction to the history
	transactions := Transactions{
		Name:       "Pembayaran",
		Price: 		totalPrice,
	}
	s.historyService.SaveTransaction(transactions)

	response := TransactionsResponse{
		Status:       "success",
		TotalPayment: totalPrice,
	}

	return response
}

