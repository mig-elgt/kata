package cyclicsort

import "testing"

func TestFinSmallestMissingPositiveNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{-3, 1, 5, 4, 2}, 3},
		{[]int{3, -2, 0, 1, 2}, 4},
		{[]int{3, 2, 5, 1}, 4},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindSmallestMissingPositiveNumber(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
