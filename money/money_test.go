package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)

	got := five.Times(2)
	want := NewDollar(10)

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = five.Times(3)
	want = NewDollar(15)

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	// Dollar
	fiveDollar := NewDollar(5)
	if fiveDollar.Equals(NewDollar(5)) != true {
		t.Errorf("(NewDollar(5) == NewDollar(5)) == false, want true")
	}

	if fiveDollar.Equals(NewDollar(6)) != false {
		t.Errorf("(NewDollar(5) == NewDollar(6)) == true, want false")
	}

	// Franc
	fiveFranc := NewFranc(5)
	if fiveFranc.Equals(NewFranc(5)) != true {
		t.Errorf("(NewFranc(5) ==  NewFranc(5)) == false, want true")
	}

	if fiveFranc.Equals(NewFranc(6)) != false {
		t.Errorf("(NewFranc(5) ==  NewFranc(6)) == true, want false")
	}

	// Dollar & Franc compare
	if fiveDollar.Equals(fiveFranc) != false {
		t.Errorf("(NewDollar(5) == NewFranc(5)) == true, want false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)

	got := five.Times(2)
	want := NewFranc(10)

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = five.Times(3)
	want = NewFranc(15)

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestCurrency(t *testing.T) {
	// Dollar
	want := "USD"
	got := NewDollar(1).currency

	if got != want {
		t.Errorf("NewDollar(1).currency == %s, want %s", got, want)
	}

	// Franc
	want = "CHF"
	got = NewFranc(1).currency

	if got != want {
		t.Errorf("NewFranc(1).currency == %s, want %s", got, want)
	}
}
