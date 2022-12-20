package grind75

import "testing"

func TestLargestRectangle(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := largestRectangleArea(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
