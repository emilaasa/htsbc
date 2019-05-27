package htscbc

import "testing"

// TestSwapValue algorithm 2.1
func TestSwapValue(t *testing.T) {
	a, b := "first", "second"
	t1 := a
	t2 := b
	b = t1
	a = t2
	// b, a = a, b // for idiomatic go

	got := b
	want := "first"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
