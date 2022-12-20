package twopointers

import "testing"

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		in   []int
		t    int
		want int
	}{
		{[]int{0, 3, 97, 102, 200}, 300, 300},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-4, -1, 1, 2}, 1, 2},
		{[]int{0, 0, 0}, 0, 0},
		{[]int{1, 1, 1, 0}, -100, 2},
	}
	for _, tc := range testCases {
		t.Run("base case: ", func(t *testing.T) {
			if got := threeSumClosest(tc.in, tc.t); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
