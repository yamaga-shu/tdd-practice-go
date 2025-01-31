package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}

	got := *five.times(2)
	want := Dollar{10}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.times(3)
	want = Dollar{15}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
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
