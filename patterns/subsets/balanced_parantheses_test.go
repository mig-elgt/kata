package subsets

import (
	"reflect"
	"testing"
)

func TestGenerateValidParentheses(t *testing.T) {
	testCases := []struct {
		in   int
		want []string
	}{
		{2, []string{"(())", "()()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := GenerateValidParentheses(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
