package bfs

import (
	"reflect"
	"testing"
)

func TestMinHeighTrees(t *testing.T) {
	testCases := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{4, [][]int{
			{1, 0},
			{1, 2},
			{1, 3},
		}, []int{1}},
		{6, [][]int{
			{3, 0},
			{3, 1},
			{3, 2},
			{3, 4},
			{5, 4},
		}, []int{3, 4}},
	}
	for _, tc := range testCases {
		if got := findMinHeightTrees(tc.n, tc.edges); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}
}
