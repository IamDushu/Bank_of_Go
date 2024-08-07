package util

// constants for all suppported currencies
const (
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, INR:
		return true
	}
	return false
}
