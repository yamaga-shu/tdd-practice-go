package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}
	product := five.times(2)

	got := product.amount
	want := 10

	if got != want {
		t.Errorf("five.times(2) == %d, want %d", got, want)
	}

	product = five.times(3)

	got = product.amount
	want = 15

	if got != want {
		t.Errorf("five.times(3) == %d, want %d", got, want)
	}
}

func TestEquals(t *testing.T) {
	five := &Dollar{5}

	if five.equals(&Dollar{5}) != true {
		t.Errorf("(&Dollar{5}).equals(&Dollar{5}) == false, want true")
	}

	if five.equals(&Dollar{6}) != false {
		t.Errorf("(&Dollar{5}).equals(&Dollar{6}) == true, want false")
	}
}
