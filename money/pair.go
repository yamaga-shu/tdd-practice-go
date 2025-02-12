package money

type Pair struct {
	from Currency // Source currency
	to   Currency // Target currency
}

// Equals checks if two Pair instances represent the same currency pair.
func (p Pair) Equals(o Pair) bool {
	return p == o
}

// HashCode returns a hash code for the Pair instance. (Currently returns 0)
// todo: implement
func (p Pair) HashCode() int {
	return 0
}
