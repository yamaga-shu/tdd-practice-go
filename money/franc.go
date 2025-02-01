package money

type Franc struct {
	money
}

// Times returns a new Franc instance multiplied by the specified rate.
func (f *Franc) Times(multiplier int) *Franc {
	return &Franc{money{f.amount * multiplier}}
}

// Equals compares whether two Franc instances have the same amount.
func (f *Franc) Equals(in *Franc) bool {
	return f.amount == in.amount
}
