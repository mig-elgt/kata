package graph

import (
	"reflect"
	"testing"
)

func TestCanConstructOriginalSequence(t *testing.T) {
	testCases := []struct {
		originalSeq []int
		seqs        [][]int
		want        bool
	}{
		{
			originalSeq: []int{1, 2, 3, 4},
			seqs: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
			},
			want: true,
		},
		{
			originalSeq: []int{1, 2, 3, 4},
			seqs: [][]int{
				{1, 2},
				{2, 3},
				{2, 4},
			},
			want: false,
		},
		{
			originalSeq: []int{3, 1, 4, 2, 5},
			seqs: [][]int{
				{3, 1, 5},
				{1, 4, 2, 5},
			},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run("test cases", func(t *testing.T) {
			if got := CanConstructOriginalSequence(tc.originalSeq, tc.seqs); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
