package grind75

import "testing"

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tc := range testCases {
		if got := majorityElement(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
