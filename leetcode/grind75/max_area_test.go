package grind75

import "testing"

func TestCalculateMaxArea(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := calculateMaxArea(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
