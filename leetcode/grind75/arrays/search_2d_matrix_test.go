package arrays

import "testing"

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		in     [][]int
		target int
		want   bool
	}{
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			13,
			false,
		},
		{
			[][]int{
				{1},
			},
			2,
			false,
		},
		{
			[][]int{
				{2},
			},
			2,
			true,
		},
		{
			[][]int{
				{1, 3},
			},
			3,
			true,
		},
		{
			[][]int{
				{1, 1},
			},
			2,
			false,
		},
	}
	for _, tc := range testCases {
		if got := searchMatrix(tc.in, tc.target); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
