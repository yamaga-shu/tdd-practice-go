package money

type Money struct {
	amount   int      // The amount of money
	currency Currency // The currency of the money
}

type Currency string

const (
	USD Currency = "USD" // US Dollar
	CHF Currency = "CHF" // Swiss Franc
)

// Equals checks if two Money instances have the same amount and currency.
func (m Money) Equals(in Money) bool {
	return m == in
}

// Times returns a new Money instance with the amount multiplied by the specified multiplier.
func (m Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus returns a new Sum instance representing the addition of the current Money instance and the addend.
func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

// Reduce converts the Money instance to the specified currency using the bank's exchange rate and returns a new Money instance.
func (m Money) Reduce(bank Bank, to Currency) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

// NewMoney creates a new Money instance with the specified amount and currency.
func NewMoney(amount int, currency Currency) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

// NewDollar creates a new Money instance with the specified amount in USD.
func NewDollar(amount int) Money {
	return NewMoney(amount, USD)
}

// NewFranc creates a new Money instance with the specified amount in CHF.
func NewFranc(amount int) Money {
	return NewMoney(amount, CHF)
}
