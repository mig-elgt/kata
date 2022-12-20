package grind75

import "testing"

func TestBracketsAreValid(t *testing.T) {
	testCases := []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"((((((()))))))", true},
		{"((((((())))))", false},
		{"((([][]}", false},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := isValid(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
