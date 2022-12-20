package binarysearch

import "testing"

func TestRotationCount(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 15, 1, 3, 8}, 2},
		{[]int{4, 5, 7, 9, 10, -1, 2}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := RotationCount(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
