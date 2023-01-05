package grind75

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want [][]int
	}{
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
	}
	for _, tc := range testCases {
		setZeroes(tc.in)
		if !reflect.DeepEqual(tc.in, tc.want) {
			t.Fatalf("got %v; want %v", tc.in, tc.want)
		}
	}
}
