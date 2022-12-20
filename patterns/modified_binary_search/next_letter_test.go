package binarysearch

import "testing"

func TestFindNextLetter(t *testing.T) {
	testCases := []struct {
		in   []rune
		key  rune
		want rune
	}{
		{[]rune{'a', 'c', 'f', 'h'}, 'f', 'h'},
		{[]rune{'a', 'c', 'f', 'h'}, 'b', 'c'},
		{[]rune{'a', 'c', 'f', 'h'}, 'm', 'a'},
		{[]rune{'a', 'c', 'f', 'h'}, 'h', 'a'},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindNextLetter(tc.in, tc.key); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
