package two_pointers

import (
	"reflect"
	"testing"
)

func TestPairWithTargetSum(t *testing.T) {
	testCases := map[string]struct {
		in     []int
		target int
		want   []int
	}{
		"base case 1": {
			in:     []int{1, 2, 3, 4, 6},
			target: 6,
			want:   []int{1, 3},
		},
		"base case 2": {
			in:     []int{1, 2, 3, 4, 6},
			target: 6,
			want:   []int{1, 3},
		},
		"not found": {
			in:     []int{1, 2, 3, 4, 6},
			target: -13,
			want:   []int{-1, -1},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := PairWithTargetSum(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
