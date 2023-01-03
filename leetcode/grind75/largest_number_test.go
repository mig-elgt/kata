package grind75

import "testing"

func TestLargestNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := largestNumber(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
