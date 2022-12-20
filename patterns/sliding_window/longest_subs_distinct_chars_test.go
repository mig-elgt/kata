package sliding_window

import "testing"

func TestLongestSubstringWithDistinctCharacters(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want int
	}{
		"base case: 3 chars": {
			in:   "aabccbb",
			want: 3,
		},
		"base case: 2 chars": {
			in:   "abbbb",
			want: 2,
		},
		"base case: 3": {
			in:   "abccde",
			want: 3,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := LongestSubstringWithDistinctCharacters(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
