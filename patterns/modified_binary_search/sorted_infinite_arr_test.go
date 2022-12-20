package binarysearch

import "testing"

func TestSearchSortedInfiniteArray(t *testing.T) {
	testCases := []struct {
		in   ArrayReader
		key  int
		want int
	}{
		{NewArrayReader([]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}), 16, 6},
		{NewArrayReader([]int{4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}), 11, -1},
		{NewArrayReader([]int{1, 3, 8, 10, 15}), 15, 4},
		{NewArrayReader([]int{1, 3, 8, 10, 15}), 200, -1},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := SearchSortedInfiniteArray(tc.in, tc.key); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
