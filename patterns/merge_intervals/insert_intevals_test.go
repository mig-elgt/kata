package mergeintervals

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		inteval   []int
		want      [][]int
	}{
		{[][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 6}, [][]int{{1, 3}, {4, 7}, {8, 12}}},
		{[][]int{{1, 3}, {5, 7}, {8, 12}}, []int{4, 10}, [][]int{{1, 3}, {4, 12}}},
		{[][]int{{2, 3}, {5, 7}}, []int{1, 4}, [][]int{{1, 4}, {5, 7}}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := InsertInterval(tc.intervals, tc.inteval); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
