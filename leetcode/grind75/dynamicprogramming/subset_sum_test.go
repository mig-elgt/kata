package dynamicprogramming

import "testing"

func TestCanPartition(t *testing.T) {
	testCases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, tc := range testCases {
		if got := canPartition(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
