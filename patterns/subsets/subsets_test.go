package subsets

import (
	"reflect"
	"testing"
)

func TestFindSubsets(t *testing.T) {
	testCases := []struct {
		in   []int
		want [][]int
	}{
		{[]int{1, 3}, [][]int{{}, {1}, {3}, {1, 3}}},
		{[]int{1, 5, 3}, [][]int{{}, {1}, {5}, {1, 5}, {3}, {1, 3}, {5, 3}, {1, 5, 3}}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindSubsets(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
