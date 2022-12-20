package graph

import "testing"

func TestIsSchedulingPossible(t *testing.T) {
	type T struct {
		tasks         int
		prerequisites [][]int
		want          bool
	}
	testCases := []T{
		{
			tasks: 3,
			prerequisites: [][]int{
				{0, 1},
				{1, 2},
			},
			want: true,
		},
		{
			tasks: 3,
			prerequisites: [][]int{
				{0, 1},
				{1, 2},
				{2, 0},
			},
			want: false,
		},
		{
			tasks: 6,
			prerequisites: [][]int{
				{2, 5},
				{0, 5},
				{0, 4},
				{1, 4},
				{3, 2},
				{1, 3},
			},
			want: true,
		},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := IsSchedulingPossible(tc.tasks, tc.prerequisites); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
