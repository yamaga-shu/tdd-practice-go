package money

type Pair struct {
	from Currency
	to   Currency
}

func (p Pair) Equals(o Pair) bool {
	return p.from == o.from && p.to == o.to
}

func (p Pair) HashCode() int {
	return 0
}
