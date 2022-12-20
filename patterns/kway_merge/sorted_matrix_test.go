package kmerge

import "testing"

func TestFindKthSmallestInSortedMatrix(t *testing.T) {
	testCases := []struct {
		in   [][]int
		k    int
		want int
	}{
		{
			in: [][]int{
				{2, 6, 8},
				{3, 7, 10},
				{5, 8, 11},
			},
			k:    5,
			want: 7,
		},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindKthSmallestInSortedMatrix(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
