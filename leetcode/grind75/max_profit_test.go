package grind75

import "testing"

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, tc := range testCases {
		if got := maxProfit(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
