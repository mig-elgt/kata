package grind75

import "testing"

func TestBasicCalculator(t *testing.T) {
	testCases := []struct {
		in   string
		want int
	}{

		{"21+1", 22},
		{"(1+1)", 2},
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"2147483647", 2147483647},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := calculate(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
