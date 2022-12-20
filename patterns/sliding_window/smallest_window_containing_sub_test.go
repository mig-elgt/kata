package sliding_window

import "testing"

func TestSmallestWindowContainingSub(t *testing.T) {
	testCases := map[string]struct {
		in      string
		pattern string
		want    string
	}{
		"base case": {
			in:      "aacbddeabcae",
			pattern: "ce",
			want:    "cae",
		},
		"repeat letters": {
			in:      "aabdec",
			pattern: "abc",
			want:    "abdec",
		},
		"unique letters": {
			in:      "aabdec",
			pattern: "abac",
			want:    "aabdec",
		},
		"base case: 1": {
			in:      "abdbca",
			pattern: "abc",
			want:    "bca",
		},
		"not found": {
			in:      "adcad",
			pattern: "abc",
			want:    "",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := SmallestWindowContainingSubstring(tc.in, tc.pattern); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
