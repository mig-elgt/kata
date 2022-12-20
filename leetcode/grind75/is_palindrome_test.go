package grind75

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := isPalindrome(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
