package currency

const (
	CLP = "CLP"
	ARS = "ARS"
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
	GBP = "GBP"
)

const commission = 0.01

var Minimum = map[string]float64{
	CLP: 100,
	ARS: 50,
	USD: 10,
	EUR: 10,
	TRY: 30,
	GBP: 40,
}

var Maximum = map[string]float64{
	CLP: 100000,
	ARS: 5000,
	USD: 10000,
	EUR: 8000,
	TRY: 4000,
	GBP: 5000,
}

var USDExchangeRate = map[string]float64{
	CLP: 882,
	ARS: 810,
	USD: 1,
	EUR: 0.9,
	TRY: 30,
	GBP: 0.8,
}

func IsValidAmount(amountToExchange float64, targetCurrency string) bool {
	return amountToExchange >= Minimum[targetCurrency] && amountToExchange <= Maximum[targetCurrency]
}

func Convert(initialCurrency string, targetCurrency string, amountToExchange float64) float64 {
	return amountToExchange / USDExchangeRate[initialCurrency] * USDExchangeRate[targetCurrency]
}

func GetWithdrawalFee(exchangedAmount float64) float64 {
	return exchangedAmount * commission
}
