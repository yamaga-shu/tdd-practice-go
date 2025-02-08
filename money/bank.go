package money

type Bank struct {
	rates []rateMap
}

type rateMap struct {
	Pair Pair
	rate int
}

func (b Bank) Reduce(source Expression, to Currency) Money {
	return source.Reduce(b, to)
}

func (b *Bank) AddRate(from Currency, to Currency, rate int) {
	b.rates = append(b.rates, rateMap{
		Pair{
			from: from,
			to:   to,
		},
		rate,
	})
}

func (b Bank) Rate(from Currency, to Currency) int {
	// the time same Currency has passed
	if from == to {
		return 1
	}

	// search from Bank.rates
	for _, rate := range b.rates {
		if rate.Pair.from == from && rate.Pair.to == to {
			return rate.rate
		}
	}

	// default rate
	return 1
}

func NewBank() Bank {
	return Bank{}
}
