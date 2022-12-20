package graph

import (
	"reflect"
	"testing"
)

func TestSortTopologicalGraph(t *testing.T) {
	type T struct {
		vertices int
		edges    [][]int
		want     []int
	}
	testCases := []T{
		{
			vertices: 4,
			edges: [][]int{
				{3, 2},
				{3, 0},
				{2, 0},
				{2, 1},
			},
			want: []int{3, 2, 0, 1},
		},
		{
			vertices: 5,
			edges: [][]int{
				{4, 2},
				{4, 3},
				{2, 0},
				{2, 1},
				{3, 1},
			},
			want: []int{4, 2, 3, 0, 1},
		},
		{
			vertices: 7,
			edges: [][]int{
				{6, 4},
				{6, 2},
				{5, 3},
				{5, 4},
				{3, 0},
				{3, 1},
				{3, 2},
				{4, 1},
			},
			want: []int{5, 6, 3, 4, 0, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := SortTopologicalGraph(tc.vertices, tc.edges...); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
