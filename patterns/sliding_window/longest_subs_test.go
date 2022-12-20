package sliding_window

import "testing"

func TestLongestSubWithMaxKDistinct(t *testing.T) {
	testCases := map[string]struct {
		k    int
		in   string
		want int
	}{
		"base case: k  = 1": {
			k:    2,
			in:   "araaci",
			want: 4,
		},
		"base case: k = 1": {
			k:    1,
			in:   "araaci",
			want: 2,
		},
		"base case: k = 3": {
			k:    3,
			in:   "cbbebi",
			want: 5,
		},
		"base case: k = 10": {
			k:    10,
			in:   "cbbebi",
			want: 6,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := LongestSubWithMaxKDistinct(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
