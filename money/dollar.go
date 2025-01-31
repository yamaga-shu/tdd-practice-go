package money

type Dollar struct {
	amount int
}

// Times returns a new Dollar instance multiplied by the specified rate.
func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

// Equals compares whether two Dollar instances have the same amount.
func (d *Dollar) Equals(in *Dollar) bool {
	return d.amount == in.amount
}
