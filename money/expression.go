package money

type Expression interface {
	Reduce(to Currency) Money
}
