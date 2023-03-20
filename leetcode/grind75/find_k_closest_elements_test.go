package grind75

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	testCases := []struct {
		in   []int
		k, x int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		// {[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run("base case: ", func(t *testing.T) {
			if got := findClosestElements(tc.in, tc.k, tc.x); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
