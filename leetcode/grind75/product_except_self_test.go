package grind75

import (
	"reflect"
	"testing"
)

func TestCalculateProductExceptSelf(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tc := range testCases {
		if got := calculateProductExceptSelf(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
