package slidingwindow

import "testing"

func TestLongestOnes(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{0, 0, 0, 1}, 4, 4},
	}
	for _, tc := range testCases {
		if got := longestOnes(tc.nums, tc.k); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
