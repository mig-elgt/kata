package grind75

import "testing"

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 987},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := myAtoi(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
