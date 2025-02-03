package money

type Money struct {
	amount int
}

// Equals compares whether two money instances have the same amount.
func (m Money) Equals(in Money) bool {
	return m == in
}
