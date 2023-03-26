package grind75

import "testing"

func TestFindMaxSum(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 5},
		{[]int{9, 8, 1, 7, 3, 4, 5}, 17},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := findMaxSum(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
