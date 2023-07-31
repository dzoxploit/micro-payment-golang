// transactionhandler_test.go
package transactions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockTransactionService struct{}

func (m *mockTransactionService) ProcessTransaction(totalAmount float64, items []Transactions) TransactionsResponse {
	// Mock the behavior of the ProcessTransaction method for testing purposes
	// Return a mock TransactionsResponse based on your test scenario
	return TransactionsResponse{
		Status:       "success",
		TotalPayment: totalAmount,
	}
}

func TestProcessTransactionHandler(t *testing.T) {
	// Create a mock TransactionService
	mockService := &mockTransactionService{}

	// Create a request body with the required JSON data
	requestBody := []byte(`{
		"total_amount": 2000.00,
		"items": [
			{"name": "Ayam Goreng", "price": 1000.00},
			{"name": "Bakso Afung", "price": 1000.00}
		],
		"credit_card": "378282246310005"
	}`)

	// Create a request with the mock request body
	req, err := http.NewRequest("POST", "/transaction/process", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create a new TransactionHandler with the mock service
	handler := NewTransactionHandler(mockService)

	// Call the handler's ProcessTransactionHandler method to execute the handler function
	handler.ProcessTransactionHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Define the expected response data for validation
	expectedResponse := `{"status":"success","total_payment":2000}`

	// Check the response body
	if strings.TrimSpace(rr.Body.String()) != expectedResponse {
        t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
    }
}