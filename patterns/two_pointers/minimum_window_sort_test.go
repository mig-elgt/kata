package two_pointers

import "testing"

func TestMinimumWindowSort(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 5, 3, 7, 10, 9, 12}, 5},
		{[]int{1, 3, 2, 0, -1, 7, 10}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{3, 2, 1}, 3},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := MinimumWindowSort(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
