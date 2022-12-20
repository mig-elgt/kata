package kmerge

import "testing"

func TestFindKthSmallest(t *testing.T) {
	testCases := []struct {
		in   [][]int
		k    int
		want int
	}{
		{[][]int{{2, 6, 8}, {3, 6, 7}, {1, 3, 4}}, 5, 4},
		{[][]int{{5, 8, 9}, {1, 7}}, 3, 7},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindKthSmallest(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
