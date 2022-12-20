package grind75

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		in   string
		want []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"", []string{}},
	}
	for _, tc := range testCases {
		if got := letterCombinations(tc.in); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
