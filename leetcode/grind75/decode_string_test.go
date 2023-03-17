package grind75

import "testing"

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"3[a]", "aaa"},
		{"10[a]", "aaaaaaaaaa"},
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tc := range testCases {
		if got := decodeString(tc.in); got != tc.want {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
