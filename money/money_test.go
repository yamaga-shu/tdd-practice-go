package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}
	product := five.Times(2)

	got := product.amount
	want := 10

	if got != want {
		t.Errorf("five.Times(2) == %d, want %d", got, want)
	}

	product = five.Times(3)

	got = product.amount
	want = 15

	if got != want {
		t.Errorf("five.Times(3) == %d, want %d", got, want)
	}
}

func TestEquals(t *testing.T) {
	five := &Dollar{5}

	if five.Equals(&Dollar{5}) != true {
		t.Errorf("(&Dollar{5}).Equals(&Dollar{5}) == false, want true")
	}

	if five.Equals(&Dollar{6}) != false {
		t.Errorf("(&Dollar{5}).Equals(&Dollar{6}) == true, want false")
	}
}
