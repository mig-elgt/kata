package grind75

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		in   []int
		want [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
			},
		},
		{
			[]int{0},
			[][]int{
				{}, {0},
			},
		},
	}
	for _, tc := range testCases {
		if got := getSubsets(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
