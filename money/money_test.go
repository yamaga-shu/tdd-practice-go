package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := &dollar{5}

	got := *five.Times(2)
	want := dollar{10}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = dollar{15}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}

func TestEquals(t *testing.T) {
	five := &dollar{5}

	if five.Equals(&dollar{5}) != true {
		t.Errorf("(&dollar{5}).equals(&dollar{5}) == false, want true")
	}

	if five.Equals(&dollar{6}) != false {
		t.Errorf("(&dollar{5}).equals(&dollar{6}) == true, want false")
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := &franc{5}

	got := *five.Times(2)
	want := franc{10}

	if got != want {
		t.Errorf("five.times(2) == %v, want %v", got, want)
	}

	got = *five.Times(3)
	want = franc{15}

	if got != want {
		t.Errorf("five.times(3) == %v, want %v", got, want)
	}
}
