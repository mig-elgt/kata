package dynamicprogramming

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := minCostClimbingStairs(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
