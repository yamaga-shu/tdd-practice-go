package money

// Expression is an interface for monetary operations, allowing for multiplication, addition, and reduction to a specific currency.
type Expression interface {
	Times(multiplier int) Expression
	Plus(addend Expression) Expression
	Reduce(bank Bank, to Currency) Money
}
