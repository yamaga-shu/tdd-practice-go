package money

type Franc struct {
	amount int
}

// Times returns a new Franc instance multiplied by the specified rate.
func (d *Franc) Times(multiplier int) *Franc {
	return &Franc{d.amount * multiplier}
}

// Equals compares whether two Franc instances have the same amount.
func (d *Franc) Equals(in *Franc) bool {
	return d.amount == in.amount
}
