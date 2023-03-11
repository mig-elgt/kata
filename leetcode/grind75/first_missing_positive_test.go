package grind75

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1}, 2},
		{[]int{0}, 1},
	}
	for _, tc := range testCases {
		t.Run("base case: ", func(t *testing.T) {
			if got := firstMissingPositive(tc.nums); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
