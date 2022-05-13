package account

type CurrencyBalance struct {
	Currency         string
	BalanceTotal     float32
	BalanceAvailable float32
	BalanceReserved  float32
}
