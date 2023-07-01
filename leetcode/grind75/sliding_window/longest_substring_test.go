package slidingwindow

import "testing"

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		// {"abcabcbb", 3},
		// {"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, tc := range testCases {
		if got := lengthOfLongestSubstring(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
