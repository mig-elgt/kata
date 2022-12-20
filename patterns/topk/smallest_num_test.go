package topk

import "testing"

func TestFindKthSmallestNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want int
	}{
		{[]int{1, 5, 12, 2, 11, 5}, 3, 5},
		{[]int{1, 5, 12, 2, 11, 5}, 4, 5},
		{[]int{5, 12, 11, -1, 12}, 3, 11},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindKthSmallestNumber(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
