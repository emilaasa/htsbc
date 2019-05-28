package htscbc

import (
	"testing"
)

// Section 2.1
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

// Section 2.2
// Given a set of n students' examination marks (in the range 0 to 100) make a
// count of the number of students that passed the examination. A pass is
// awarded for all marks of 50 and above.
func TestCounting2_2(t *testing.T) {
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

// Modify algorithm so marks are read until an end-of-file is encountered.
// For this set of marks determine the total number of marks, passes,
// and percentage pass rate
func Test2_2_1(t *testing.T) {
	// Maybe create a temporary file?
	// Write some data to the temp file
	// Read in the same data (lol)
	// count stuff
	// Let's just start with the data pass
	testData := []struct {
		cutoff  int
		total   int
		passing int
		input   []int
	}{
		{
			cutoff:  50,
			total:   4,
			passing: 2,
			input:   []int{51, 49, 50, 1},
		},
	}

	for _, v := range testData {
		passing, total, _ := countAbove2(v.cutoff, v.input)
		if passing != v.passing {
			t.Errorf("%d %d", passing, v.passing)
		}

		if total != v.total {
			t.Errorf("%d %d", total, v.total)
		}

	}
}

func countAbove2(cutoff int, d []int) (passing int, total int, percentPassing float64) {
	for _, v := range d {
		total++
		if v >= cutoff {
			passing++
		}
	}
	percentPassing = float64(passing / total)
	return
}
