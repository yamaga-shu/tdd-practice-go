package money

type Bank struct {
}

func (b Bank) Reduce(source Expression, to Currency) Money {
	return source.Reduce(b, to)
}

func (b Bank) AddRate(from Currency, to Currency, rate int) {

}

func (b Bank) Rate(from Currency, to Currency) int {
	switch from {
	case USD:
		return 1
	case CHF:
		return 2
	default:
		return 1
	}
}

func NewBank() Bank {
	return Bank{}
}
