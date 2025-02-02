package money

type money struct {
	amount int
}

// Equals compares whether two money instances have the same amount.
func (m money) Equals(in money) bool {
	return m.amount == in.amount
}
