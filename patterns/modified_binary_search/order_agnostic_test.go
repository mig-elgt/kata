package binarysearch

import "testing"

func TestOrderAgnosticBinarySearch(t *testing.T) {
	testCases := []struct {
		nums []int
		key  int
		want int
	}{
		{[]int{4, 6, 10}, 10, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 5, 4},
		{[]int{10, 6, 4}, 10, 0},
		{[]int{10, 6, 4}, 4, 2},
		{[]int{10, 6, 4}, 8, -1},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := OrderAgnosticBinarySearch(tc.nums, tc.key); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
