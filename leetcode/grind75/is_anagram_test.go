package grind75

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := isAnagram(tc.s, tc.t); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
