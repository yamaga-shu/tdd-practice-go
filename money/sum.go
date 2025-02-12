package money

type Sum struct {
	augend Expression // 被加算数 (the first addend)
	addend Expression // 加数 (the second addend)
}

// Times returns a new Sum instance with each addend multiplied by the specified multiplier.
func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}

// Plus returns a new Sum instance representing the addition of the current Sum instance and the addend.
func (s Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

// Reduce converts the Sum instance to the specified currency using the bank's exchange rate and returns a new Money instance.
func (s Sum) Reduce(bank Bank, to Currency) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}

// NewSum creates a new Sum instance with the specified augend and addend.
func NewSum(augend Expression, addend Expression) Sum {
	return Sum{
		augend: augend,
		addend: addend,
	}
}
