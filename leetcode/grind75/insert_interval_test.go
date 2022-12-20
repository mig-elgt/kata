package grind75

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	testCases := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			intervals: [][]int{
				{1, 3},
				{6, 9},
			},
			newInterval: []int{2, 5},
			want: [][]int{
				{1, 5},
				{6, 9},
			},
		},
		{
			intervals: [][]int{
				{1, 2},
				{3, 5},
				{6, 7},
				{8, 10},
				{12, 16},
			},
			newInterval: []int{4, 8},
			want: [][]int{
				{1, 2},
				{3, 10},
				{12, 16},
			},
		},
	}
	for _, tc := range testCases {
		if got := insertInteval(tc.intervals, tc.newInterval); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
