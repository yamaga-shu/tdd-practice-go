package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}

	got := five.Times(2).amount
	want := 10

	if got != want {
		t.Errorf("add() = %v, want %v", got, want)
	}
}
