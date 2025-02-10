package money

type Sum struct {
	augend Expression // 被加算数
	addend Expression // 加数
}

func (s Sum) Times(multiplier int) Expression {
	return NewSum(s.augend.Times(multiplier), s.addend.Times(multiplier))
}

func (s Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

func (s Sum) Reduce(bank Bank, to Currency) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount

	return NewMoney(amount, to)
}

func NewSum(augend Expression, addend Expression) Sum {
	return Sum{
		augend: augend,
		addend: addend,
	}
}
