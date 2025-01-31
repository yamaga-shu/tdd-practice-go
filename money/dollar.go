package money

type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

func (d *Dollar) equals(in *Dollar) bool {
	return d.amount == in.amount
}
