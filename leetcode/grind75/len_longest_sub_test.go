package grind75

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"au", 2},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := lengthOfLongestSubstring(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
