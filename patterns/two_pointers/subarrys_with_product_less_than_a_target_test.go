package two_pointers

import (
	"reflect"
	"testing"
)

func TestSubArraysWithProductLessThanATarget(t *testing.T) {
	testCases := map[string]struct {
		in     []int
		target int
		want   [][]int
	}{
		"base case 1": {
			in:     []int{2, 5, 3, 10},
			target: 30,
			want: [][]int{
				{2},
				{2, 5},
				{5},
				{5, 3},
				{3},
				{10},
			},
		},
		"base case 2": {
			in:     []int{8, 2, 6, 5},
			target: 50,
			want: [][]int{
				{8},
				{8, 2},
				{2},
				{2, 6},
				{6},
				{6, 5},
				{5},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := SubArraysWithProductLessThanATarget(tc.in, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
