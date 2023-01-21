package grind75

import (
	"reflect"
	"testing"
)

func TestFindMinHeightTrees(t *testing.T) {
	testCases := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{4, [][]int{{1, 0}, {1, 2}, {1, 3}}, []int{1}},
		{6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}, []int{3, 4}},
		{1, [][]int{}, []int{0}},
		{2, [][]int{{0, 1}}, []int{1, 0}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := findMinHeightTrees(tc.n, tc.edges); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
