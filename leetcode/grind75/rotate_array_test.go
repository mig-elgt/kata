package grind75

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if rotate(tc.in, tc.k); !reflect.DeepEqual(tc.in, tc.want) {
				t.Fatalf("got %v; want %v", tc.in, tc.want)
			}
		})
	}
}
