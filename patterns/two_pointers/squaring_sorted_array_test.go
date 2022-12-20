package two_pointers

import (
	"reflect"
	"testing"
)

func TestSquaringSortedArray(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		want []int
	}{
		"base case 1": {
			in:   []int{-2, -1, 0, 2, 3},
			want: []int{0, 1, 4, 4, 9},
		},
		"base case 2": {
			in:   []int{-3, -1, 0, 1, 2},
			want: []int{0, 1, 1, 4, 9},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := SquaringSortedArray(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
