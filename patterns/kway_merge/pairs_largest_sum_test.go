package kmerge

import (
	"reflect"
	"testing"
)

func TestFindKLargestPairs(t *testing.T) {
	testCases := []struct {
		l1   []int
		l2   []int
		k    int
		want [][]int
	}{
		{
			l1: []int{9, 8, 2},
			l2: []int{6, 3, 1},
			k:  3,
			want: [][]int{
				{9, 3},
				{8, 6},
				{9, 6},
			},
		},
		{
			l1: []int{5, 2, 1},
			l2: []int{2, -1},
			k:  3,
			want: [][]int{
				{5, -1},
				{2, 2},
				{5, 2},
			},
		},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindKLargestPairs(tc.l1, tc.l2, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
