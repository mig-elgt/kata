package grind75

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		in   []int
		want [][]int
	}{
		{
			in: []int{1, 2, 3},
			want: [][]int{
				{3, 2, 1},
				{2, 3, 1},
				{2, 1, 3},
				{3, 1, 2},
				{1, 3, 2},
				{1, 2, 3},
			},
		},
		{
			in: []int{5, 4, 6, 2},
			want: [][]int{
				{2, 6, 4, 5},
				{6, 2, 4, 5},
				{6, 4, 2, 5},
				{6, 4, 5, 2},
				{2, 4, 6, 5},
				{4, 2, 6, 5},
				{4, 6, 2, 5},
				{4, 6, 5, 2},
				{2, 4, 5, 6},
				{4, 2, 5, 6},
				{4, 5, 2, 6},
				{4, 5, 6, 2},
				{2, 6, 5, 4},
				{6, 2, 5, 4},
				{6, 5, 2, 4},
				{6, 5, 4, 2},
				{2, 5, 6, 4},
				{5, 2, 6, 4},
				{5, 6, 2, 4},
				{5, 6, 4, 2},
				{2, 5, 4, 6},
				{5, 2, 4, 6},
				{5, 4, 2, 6},
				{5, 4, 6, 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := permute(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
