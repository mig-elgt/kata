package grind75

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	}
	for _, tc := range testCases {
		if got := mergeIntervals(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
