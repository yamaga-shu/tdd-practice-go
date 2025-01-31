package money

type Dollar struct {
	amount int
}

func (d *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{d.amount * multiplier}
}

func (d *Dollar) Equals(in *Dollar) bool {
	return d.amount == in.amount
}
