package money

type dollar struct {
	amount int
}

// Times returns a new dollar instance multiplied by the specified rate.
func (d *dollar) Times(multiplier int) *dollar {
	return &dollar{d.amount * multiplier}
}

// Equals compares whether two dollar instances have the same amount.
func (d *dollar) Equals(in *dollar) bool {
	return d.amount == in.amount
}
