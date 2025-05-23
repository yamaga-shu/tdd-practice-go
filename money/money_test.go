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
	want := USD
	got := NewDollar(1).currency

	if got != want {
		t.Errorf("NewDollar(1).currency == %s, want %s", got, want)
	}

	// Franc
	want = CHF
	got = NewFranc(1).currency

	if got != want {
		t.Errorf("NewFranc(1).currency == %s, want %s", got, want)
	}
}

func TestSimpleAddition(t *testing.T) {
	var five Expression = NewDollar(5)
	var sum Expression = five.Plus(five)
	var bank Bank = NewBank()

	var got Money = bank.Reduce(sum, USD)

	want := NewDollar(10)

	if got != want {
		t.Errorf("NewDollar(5).Plus(NewDollar(5)) == %v, want %v", got, want)
	}
}

func TestRuduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	var bank Bank = NewBank()

	var got Money = bank.Reduce(sum, USD)

	want := NewDollar(7)

	if got != want {
		t.Errorf("bank.Reduce(sum, USD) == %v, want %v", got, want)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	got := bank.Reduce(NewDollar(1), USD)

	want := NewDollar(1)

	if got != want {
		t.Errorf("bank.Reduce(NewDollar(1), USD) == %v, want %v", got, want)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate(CHF, USD, 2)
	got := bank.Reduce(NewFranc(2), USD)

	want := NewDollar(1)

	if got != want {
		t.Errorf("bank.Reduce(NewFranc(2), USD) == %v, want %v", got, want)
	}
}

func TestIdentityRate(t *testing.T) {
	bank := NewBank()
	got := bank.Rate(USD, USD)

	want := 1

	if got != want {
		t.Errorf("bank.Rate(USD, USD) == %v, want %v", got, want)
	}
}

func TestMixedAddition(t *testing.T) {
	var fiveDollar Expression = NewDollar(5)
	var tenFranc Expression = NewFranc(10)
	bank := NewBank()
	bank.AddRate(CHF, USD, 2)

	got := bank.Reduce(fiveDollar.Plus(tenFranc), USD)

	want := NewDollar(10)

	if got != want {
		t.Errorf("bank.Reduce(fiveDollar.Plus(tenFranc), USD) == %v, want %v", got, want)
	}
}

func TestSumPlusMoney(t *testing.T) {
	var fiveDollar Expression = NewDollar(5)
	var tenFranc Expression = NewFranc(10)
	bank := NewBank()
	bank.AddRate(CHF, USD, 2)

	var sum Expression = NewSum(fiveDollar, tenFranc).Plus(fiveDollar)

	got := bank.Reduce(sum, USD)

	want := NewDollar(15)

	if got != want {
		t.Errorf("bank.Reduce(sum, USD) == %v, want %v", got, want)
	}
}

func TestSumTimes(t *testing.T) {
	var fiveDollar Expression = NewDollar(5)
	var tenFranc Expression = NewFranc(10)
	var bank Bank = NewBank()
	bank.AddRate(CHF, USD, 2)
	var sum Expression = NewSum(fiveDollar, tenFranc).Times(2)

	got := bank.Reduce(sum, USD)

	want := NewDollar(20)

	if got != want {
		t.Errorf("bank.Reduce(sum, USD) == %v, want %v", got, want)
	}
}
