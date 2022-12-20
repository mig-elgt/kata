package grind75

import "testing"

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		s, t, want string
	}{
		{"acbbaca", "aba", "baca"},
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "b", "b"},
		{"cabefgecdaecf", "cae", "aec"},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := minWindow(tc.s, tc.t); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
