package sliding_window

import "testing"

func TestPermutationInString(t *testing.T) {
	testCases := map[string]struct {
		in      string
		pattern string
		want    bool
	}{
		"has pattern": {
			in:      "oidbcaf",
			pattern: "abc",
			want:    true,
		},
		"has no pattern": {
			in:      "odicf",
			pattern: "dc",
			want:    false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := PermutationInString(tc.in, tc.pattern); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
