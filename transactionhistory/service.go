package transactionhistory

import (
	"log"

	"github.com/dzoxploit/micro-payment-golang/database"
	"github.com/dzoxploit/micro-payment-golang/transaction"
)

type TransactionHistoryService interface {
	SaveTransaction(transaction transaction.Transaction)
}

type transactionHistoryService struct{}

func NewTransactionHistoryService() TransactionHistoryService {
	return &transactionHistoryService{}
}

func (s *transactionHistoryService) SaveTransaction(transaction transaction.Transaction) {
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
