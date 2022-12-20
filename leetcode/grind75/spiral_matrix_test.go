package grind75

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want []int
	}{
		{
			in: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			in: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			in: [][]int{
				{3},
				{2},
			},
			want: []int{3, 2},
		},
	}
	for _, tc := range testCases {
		if got := spiralOrder(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
