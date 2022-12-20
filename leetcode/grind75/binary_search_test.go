package grind75

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		in     []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, 13, -1},
		{[]int{-1, 0, 5}, -1, 0},
	}
	for _, tc := range testCases {
		if got := search(tc.in, tc.target); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
