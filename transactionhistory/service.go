package transactionhistory
import ()

type TransactionHistoryService interface {
	SaveTransaction(transaction Transaction)
}

type transactionHistoryService struct{}

func NewTransactionHistoryService() TransactionHistoryService {
	return &transactionHistoryService{}
}

func (s *transactionHistoryService) SaveTransaction(transaction Transaction) {
	// Save the transaction to the database or any other data store
	// ...
}
