package money

type Franc struct {
	Money
}

// Times returns a new Franc instance multiplied by the specified rate.
func (f Franc) Times(multiplier int) Franc {
	return NewFranc(f.amount * multiplier)
}
