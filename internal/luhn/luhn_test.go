package luhn

import "testing"

func TestValidateCard(t *testing.T) {
	testCases := []struct {
		cardName   string
		cardNumber string
		expected   bool
	}{
		{"MasterCard", "5555555555554444", true},
		{"MasterCard", "5105105105105100", true},
		{"Visa", "4111111111111111", true},
		{"Visa", "4012888888881881", true},
	}

	for _, tc := range testCases {
		t.Run(tc.cardName, func(t *testing.T) {
			result := ValidateCard(tc.cardNumber)
			if result != tc.expected {
				t.Errorf("For %s expected %v, got %v", tc.cardNumber, tc.expected, result)
			}
		})
	}
}
