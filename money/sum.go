package money

type Sum struct {
	augend Money
	addend Money
}

func (s Sum) Reduce(bank Bank, to Currency) Money {
	amount := s.augend.amount + s.addend.amount

	return NewMoney(amount, to)
}
