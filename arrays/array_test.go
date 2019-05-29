package array

import (
	"reflect"
	"testing"
)

func TestReverseOrderOfArray(t *testing.T) {
	cases := []struct {
		want  []int
		input []int
	}{
		{want: []int{3, 2, 1}, input: []int{1, 2, 3}},
		{want: []int{4, 3, 2, 1}, input: []int{1, 2, 3, 4}},
	}

	for _, tc := range cases {
		got := reverse(tc.input)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("got: %v, want: %v", got, tc.want)
		}
	}
}

func reverse(s []int) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[len(s)-1-i] = v
	}
	return r
}

func TestHistogram(t *testing.T) {
	// Represents the laziest way to write up some scores:
	// 2 ppl with 0 score, 2 with 1
	want := make([]int, 101)
	want[0] = 2
	want[1] = 2
	input := []int{0, 0, 1, 1}

	got := countScores(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

}

func countScores(s []int) []int {
	r := make([]int, 101)
	for _, v := range s {
		r[v]++
	}
	return r
}

func TestFindMaximum(t *testing.T) {
	want := 5
	in := []int{1, 2, 3, 5, 5}
	got := findMax(in)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func findMax(xs []int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func TestMinimum(t *testing.T) {
	var tests = []struct {
		name     string
		expected int
		given    []int
	}{
		{"simple case", 1, []int{2, 1, 3}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual := findMin(tt.given)
			if actual != tt.expected {
				t.Errorf("(%v): expected %v, actual %v", tt.given, tt.expected, actual)
			}

		})
	}

}

func findMin(xs []int) int {
	r := xs[0]
	for _, x := range xs {
		if x < r {
			r = x
		}
	}
	return r
}
