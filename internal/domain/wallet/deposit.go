package wallet

type Deposit struct {
	money      Money
	equivalent Equivalent
}

func RecreateDeposit(deposit Money, equivalent Equivalent) Deposit {
	return Deposit{
		money:      deposit,
		equivalent: equivalent,
	}
}

func (deposit Deposit) Money() Money {
	return deposit.money
}

func (deposit Deposit) Equivalent() Equivalent {
	return deposit.equivalent
}
