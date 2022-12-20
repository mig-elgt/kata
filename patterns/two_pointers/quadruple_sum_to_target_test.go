package two_pointers

import (
	"reflect"
	"testing"
)

func TestQuadupleSumToTarget(t *testing.T) {
	testCases := map[string]struct {
		in     []int
		target int
		want   [][]int
	}{
		"base case 1": {
			in:     []int{4, 1, 2, -1, 1, -3},
			target: 1,
			want: [][]int{
				{-3, -1, 1, 4},
				{-3, 1, 1, 2},
			},
		},
		"base case 2": {
			in:     []int{2, 0, -1, 1, -2, 2},
			target: 2,
			want: [][]int{
				{-2, 0, 2, 2},
				{-1, 0, 1, 2},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := QuadrupleSumToTarget(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
