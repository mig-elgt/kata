package arrays

import "testing"

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{-1, -1, 1}, 0, 1},
		// {[]int{1, -1, 0}, 0, 3},
	}
	for _, tc := range testCases {
		if got := subarraySum(tc.in, tc.k); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
