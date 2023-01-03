package grind75

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"foobar", "foobar", "foobar"}, "foobar"},
		{[]string{"ab", "a"}, "a"},
	}
	for _, tc := range testCases {
		if got := longestCommonPrefix(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
