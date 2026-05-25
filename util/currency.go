package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	JPY = "JPY"
	KRW = "KRW"
	CNY = "CNY"
	INR = "INR"
	BRL = "BRL"
	MXN = "MXN"
	CAD = "CAD"
	AUD = "AUD"
	ZMW = "ZMW"
)

// IsSupportedCurrency checks if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, JPY, KRW, CNY, INR, BRL, MXN, CAD, AUD, ZMW:
		return true
	}
	return false
}
