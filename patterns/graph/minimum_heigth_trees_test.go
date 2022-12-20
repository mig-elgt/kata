package graph

import (
	"reflect"
	"testing"
)

func TestFindMimmumHeightTreeNodes(t *testing.T) {
	testCases := []struct {
		vertices int
		edges    [][]int
		want     []int
	}{
		{5, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 2}},
		{4, [][]int{{0, 1}, {0, 2}, {2, 3}}, []int{0, 2}},
		{4, [][]int{{0, 1}, {1, 2}, {1, 3}}, []int{1}},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindMimmumHeightTreeNodes(tc.vertices, tc.edges); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
