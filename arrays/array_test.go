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
