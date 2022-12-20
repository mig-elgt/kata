package grind75

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := containsDuplicate(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
