package grind75

import "testing"

func TestRansomNote(t *testing.T) {
	testCases := []struct {
		ransomNote, magazine string
		want                 bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, tc := range testCases {
		if got := canConstruct(tc.ransomNote, tc.magazine); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
