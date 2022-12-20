package xor

import "testing"

func TestFindMissingNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 5, 2, 6, 4}, 3},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindMissingNumber(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
