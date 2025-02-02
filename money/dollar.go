package money

type Dollar struct {
	money
}

// NewDollar returns a new Dollar instance have the specified amount.
func NewDollar(amount int) Dollar {
	return Dollar{money{amount}}
}

// Times returns a new Dollar instance multiplied by the specified rate.
func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}

// Equals compares whether two Dollar instances have the same amount.
func (d Dollar) Equals(in Dollar) bool {
	return d.amount == in.amount
}
