package subsets

import (
	"reflect"
	"testing"
)

func TestGetPermutations(t *testing.T) {
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
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := GetPermutations(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
