package mergeintervals

import (
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	testCases := []struct {
		arr1 [][]int
		arr2 [][]int
		want [][]int
	}{
		{[][]int{{1, 3}, {5, 6}, {7, 9}}, [][]int{{2, 3}, {5, 7}}, [][]int{{2, 3}, {5, 6}, {7, 7}}},
		{[][]int{{1, 3}, {5, 7}, {9, 12}}, [][]int{{5, 10}}, [][]int{{5, 7}, {9, 10}}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := IntervalIntersection(tc.arr1, tc.arr2); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
