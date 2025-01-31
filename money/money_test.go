package money

import "testing"

func TestMultiplication(t *testing.T) {
	cases := []struct {
		in, multiplier, want int
	}{
		{2, 5, 10},
	}

	for _, c := range cases {
		got := (&Dollar{c.in}).Times(c.multiplier).amount
		if got != c.want {
			t.Errorf("(&Dollar{%d}).Times(%d).amount == %d, want %d", c.in, c.multiplier, got, c.want)
		}
	}
}
