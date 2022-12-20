package grind75

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		in   []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tc := range testCases {
		if got := threeSum(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
