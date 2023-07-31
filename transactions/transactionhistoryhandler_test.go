// transactionhistoryhandler_test.go
package transactions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTransactionHistoryService struct{}

func (m *mockTransactionHistoryService) SaveTransaction(transaction Transactions) {
	// Mock the behavior of the SaveTransaction method for testing purposes
	// You can perform additional checks or actions based on your test scenario
}

func TestSaveTransactionHandler(t *testing.T) {
	// Create a mock TransactionHistoryService
	mockService := &mockTransactionHistoryService{}

	// Create a request body with the required JSON data
	requestBody := []byte(`{
		"transaction": {
			"name": "Product A",
			"price": 1000.00
		}
	}`)

	// Create a request with the mock request body
	req, err := http.NewRequest("POST", "/transactionhistory/save", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create a new TransactionHistoryHandler with the mock service
	handler := NewTransactionHistoryHandler(mockService)

	// Call the handler's SaveTransactionHandler method to execute the handler function
	handler.SaveTransactionHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
