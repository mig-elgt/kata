package two_pointers

import (
	"reflect"
	"testing"
)

func TestDutchNationalFlagProblem(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		want []int
	}{
		"base case 1": {
			in:   []int{1, 0, 2, 1, 0},
			want: []int{0, 0, 1, 1, 2},
		},
		"base case 2": {
			in:   []int{2, 2, 0, 1, 2, 0},
			want: []int{0, 0, 1, 2, 2, 2},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := DutchNationalFlagProblem(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
