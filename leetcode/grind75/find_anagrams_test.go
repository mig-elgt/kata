package grind75

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		s, p string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"baa", "aa", []int{1}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := findAnagrams(tc.s, tc.p); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
