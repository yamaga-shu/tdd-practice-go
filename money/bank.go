package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to Currency) Money {
	return source.Reduce(to)
}

func NewBank() Bank {
	return Bank{}
}
