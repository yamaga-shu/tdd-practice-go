package money

type Bank struct {
	rates []rateMap // A collection of exchange rates
}

type rateMap struct {
	Pair Pair // Currency pair
	rate int  // Exchange rate
}

// Reduce converts the source Expression to the specified currency using the bank's exchange rate and returns a new Money instance.
func (b Bank) Reduce(source Expression, to Currency) Money {
	return source.Reduce(b, to)
}

// AddRate adds a new exchange rate to the bank for the specified currency pair.
func (b *Bank) AddRate(from Currency, to Currency, rate int) {
	b.rates = append(b.rates, rateMap{
		Pair{
			from: from,
			to:   to,
		},
		rate,
	})
}

// Rate returns the exchange rate for the specified currency pair. If the currencies are the same, returns 1.
func (b Bank) Rate(from Currency, to Currency) int {
	if from == to {
		return 1
	}

	for _, rate := range b.rates {
		if rate.Pair.from == from && rate.Pair.to == to {
			return rate.rate
		}
	}

	return 1
}

// NewBank creates a new Bank instance with no exchange rates.
func NewBank() Bank {
	return Bank{}
}
