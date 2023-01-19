package grind75

import "testing"

func TestOrangesRotting(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}, 4},
		{[][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{1, 2}}, 1},
	}
	for _, tc := range testCases {
		t.Run("base case: ", func(t *testing.T) {
			if got := orangesRotting(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
