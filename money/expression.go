package money

type Expression interface {
	Reduce(bank Bank, to Currency) Money
}
