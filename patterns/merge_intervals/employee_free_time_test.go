package mergeintervals

import (
	"reflect"
	"testing"
)

func TestFindEmployeeFreeTime(t *testing.T) {
	testCases := []struct {
		in   [][]int
		want [][]int
	}{
		{[][]int{{1, 3}, {5, 6}, {2, 3}, {6, 8}}, [][]int{{3, 5}}},
		{[][]int{{1, 3}, {9, 12}, {2, 4}, {6, 8}}, [][]int{{4, 6}, {8, 9}}},
		{[][]int{{1, 3}, {2, 4}, {3, 5}, {7, 9}}, [][]int{{5, 7}}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindEmployeeFreeTime(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
