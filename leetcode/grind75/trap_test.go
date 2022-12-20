package grind75

import "testing"

func TestTrap(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{4, 2, 3}, 1},
	}
	for _, tc := range testCases {
		if got := trap(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
