package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &Dollar{money{5}}

	got := *five.Times(2)
	want := Dollar{money{10}}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = Dollar{money{15}}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	five := &Dollar{money{5}}

	if five.Equals(&Dollar{money{5}}) != true {
		t.Errorf("(&Dollar{money{5}}).equals(&Dollar{money{5}}) == false, want true")
	}

	if five.Equals(&Dollar{money{6}}) != false {
		t.Errorf("(&Dollar{money{5}}).equals(&Dollar{money{6}}) == true, want false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := &Franc{money{5}}

	got := *five.Times(2)
	want := Franc{money{10}}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = Franc{money{15}}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}
