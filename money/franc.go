package money

type franc struct {
	amount int
}

// Times returns a new franc instance multiplied by the specified rate.
func (d *franc) Times(multiplier int) *franc {
	return &franc{d.amount * multiplier}
}

// Equals compares whether two franc instances have the same amount.
func (d *franc) Equals(in *franc) bool {
	return d.amount == in.amount
}
