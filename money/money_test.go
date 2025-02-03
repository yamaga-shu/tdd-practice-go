package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{Money{5}}

	got := five.Times(2)
	want := Dollar{Money{10}}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = five.Times(3)
	want = Dollar{Money{15}}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	// Dollar
	fiveDollar := Dollar{Money{5}}
	if fiveDollar.Equals(Dollar{Money{5}}.Money) != true {
		t.Errorf("(Dollar{Money{5}}.Money == Dollar{Money{5}}.Money) == false, want true")
	}

	if fiveDollar.Equals(Dollar{Money{6}}.Money) != false {
		t.Errorf("(Dollar{Money{5}}.Money == Dollar{Money{6}}.Money) == true, want false")
	}

	// Franc
	fiveFranc := Franc{Money{5}}
	if fiveFranc.Equals(Franc{Money{5}}.Money) != true {
		t.Errorf("(Franc{Money{5}}.Money == Franc{Money{5}}.Money) == false, want true")
	}

	if fiveFranc.Equals(Franc{Money{6}}.Money) != false {
		t.Errorf("(Franc{Money{5}}.Money == Franc{Money{6}}.Money) == true, want false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := Franc{Money{5}}

	got := five.Times(2)
	want := Franc{Money{10}}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = five.Times(3)
	want = Franc{Money{15}}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}
