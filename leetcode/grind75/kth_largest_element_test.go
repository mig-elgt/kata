package grind75

import "testing"

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tc := range testCases {
		if got := findKthLargest(tc.in, tc.k); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
