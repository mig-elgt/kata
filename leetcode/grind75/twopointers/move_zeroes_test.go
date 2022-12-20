package twopointers

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{0, 1, 2}, []int{1, 2, 0}},
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 2, 0, 3}, []int{1, 2, 3, 0}},
	}
	for _, tc := range testCases {
		if moveZeroes(tc.in); !reflect.DeepEqual(tc.in, tc.want) {
			t.Fatalf("got %v; want %v", tc.in, tc.want)
		}
	}
}
