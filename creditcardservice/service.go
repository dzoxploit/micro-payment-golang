package creditcardservice

import (
	"strconv"
	"strings"
)

type CreditCardService interface {
	ValidateCreditCard(creditCardNumber string) bool
}

type creditCardService struct{}

func NewCreditCardService() CreditCardService {
	return &creditCardService{}
}

func (s *creditCardService) ValidateCreditCard(creditCardNumber string) bool {
	// Remove any spaces or dashes from the credit card number
	creditCardNumber = strings.ReplaceAll(creditCardNumber, " ", "")
	creditCardNumber = strings.ReplaceAll(creditCardNumber, "-", "")

	// Check if the credit card number contains only digits
	if _, err := strconv.Atoi(creditCardNumber); err != nil {
		return false
	}

	// The Luhn algorithm validation
	var sum int
	alternate := false
	for i := len(creditCardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(creditCardNumber[i]))
		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}