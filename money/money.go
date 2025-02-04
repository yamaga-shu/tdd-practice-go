package money

type Money struct {
	amount   int
	currency string
}

// Equals compares whether two money instances have the same amount.
func (m Money) Equals(in Money) bool {
	return m == in
}

// NewDollar returns a new Dollar instance have the specified amount.
func NewDollar(amount int) Dollar {
	return Dollar{
		Money{
			amount:   amount,
			currency: "USD",
		},
	}
}

// NewFranc returns a new Franc instance have the specified amount.
func NewFranc(amount int) Franc {
	return Franc{
		Money{
			amount:   amount,
			currency: "CHF",
		},
	}
}
