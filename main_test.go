package htscbc

import (
	"testing"
)

// TestSwapValue algorithm 2.1
func TestSwapValue(t *testing.T) {
	a, b := "first", "second"
	temp := a
	a = b
	b = temp
	// b, a = a, b // for idiomatic go

	got := b
	want := "first"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}

	got = a
	want = "second"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestSwapValue211(t *testing.T) {
	a, b := "raspberry", "lemonade"
	b, a = a, b

	if a != "lemonade" {
		t.Errorf("got: %s, want: %s", a, "lemonade")
	}

	if b != "raspberry" {
		t.Errorf("got: %s, want: %s", b, "raspberry")
	}
}

func TestSwapValue212(t *testing.T) {
	// Exchange a -> b -> c -> a

	a, b, c := "a", "b", "c"

	t1 := a
	a = c
	c = b
	b = t1

	got := a
	want := "c"
	if got != want {
		t.Errorf("got: %s, want: %s", a, "c")
	}

	got = b
	want = "a"
	if got != want {
		t.Errorf("got: %s, want: %s", b, "a")
	}
}

func TestSwapValue213(t *testing.T) {
	// Shift left and keep the last element
	a := []string{"a", "b", "c", "d"}
	b := append(a[1:len(a)], a[0])

	want := []string{"b", "c", "d", "a"}
	got := b

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("got: %s, want: %s", got, want)
		}
	}

}

// Given two variables of integer type a and b, exhange their values without
// using a third temporary value
func TestSwapValue214(t *testing.T) {
	a, b := 5, 10
	b = b - a
	a = a + a
	if a != 10 {
		t.Errorf("got: %d, want: %d", a, 10)
	}
	if b != 5 {
		t.Errorf("got: %d, want: %d", b, 5)

	}
}

// Given a set of n students' examination marks (in the range 0 to 100) make a
// count of the number of students that passed the examination. A pass is
// awarded for all marks of 50 and above.
func TestCounting22(t *testing.T) {
	studentMarks := []int{1, 100, 50, 51, 0, 49}
	want := 3 // 3 students passed

	got := countAbove(50, studentMarks)

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func countAbove(cutoff int, d []int) int {
	var count int
	for _, v := range d {
		if v >= cutoff {
			count++
		}
	}
	return count
}
