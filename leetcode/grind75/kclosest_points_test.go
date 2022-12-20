package grind75

import (
	"reflect"
	"testing"
)

func TestKclosestPoints(t *testing.T) {
	testCases := []struct {
		in   [][]int
		k    int
		want [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{-2, 4}, {3, 3}}},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := kClosest(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
