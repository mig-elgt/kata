package arrays

import "testing"

func TestMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		a    []int
		b    []int
		want float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, tc := range testCases {
		if got := findMedianSortedArrays(tc.a, tc.b); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
