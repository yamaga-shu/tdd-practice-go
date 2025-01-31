package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{5}

	got := *five.Times(2)
	want := Dollar{10}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = Dollar{15}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	five := &Dollar{5}

	if five.Equals(&Dollar{5}) != true {
		t.Errorf("(&Dollar{5}).equals(&Dollar{5}) == false, want true")
	}

	if five.Equals(&Dollar{6}) != false {
		t.Errorf("(&Dollar{5}).equals(&Dollar{6}) == true, want false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := &Franc{5}

	got := *five.Times(2)
	want := Franc{10}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = Franc{15}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}
