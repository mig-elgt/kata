package topk

import "testing"

func TestFindSumOfElements(t *testing.T) {
	testCases := []struct {
		in     []int
		k1, k2 int
		want   int
	}{
		{[]int{1, 3, 12, 5, 15, 11}, 3, 6, 23},
		{[]int{3, 5, 8, 7}, 1, 4, 12},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindSumOfElements(tc.in, tc.k1, tc.k2); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
