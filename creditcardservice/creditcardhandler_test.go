// creditcardhandler_test.go
package creditcardservice

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockCreditCardService struct{}

func (m *mockCreditCardService) ValidateCreditCard(creditCardNumber string) bool {
	// Mock the behavior of the ValidateCreditCard method for testing purposes
	// Return true or false based on your test scenario
	return true
}

func TestValidateCreditCardHandler(t *testing.T) {
	// Create a mock CreditCardService
	mockService := &mockCreditCardService{}

	// Create a request body with the required JSON data
	requestBody := []byte(`{"credit_card_number": "1234567890123456"}`)

	// Create a request with the mock request body
	req, err := http.NewRequest("POST", "/creditcard/validate", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Create a new CreditCardHandler with the mock service
	handler := NewCreditCardHandler(mockService)

	// Call the handler's ValidateCreditCardHandler method to execute the handler function
	handler.ValidateCreditCardHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Define the expected response data for validation
	// Define the expected response data for validation
    expectedResponse := `{"is_valid":true}`

    // Check the response body after trimming white spaces
    if strings.TrimSpace(rr.Body.String()) != expectedResponse {
        t.Errorf("Handler returned unexpected body: got %v, want %v", rr.Body.String(), expectedResponse)
    }
}
