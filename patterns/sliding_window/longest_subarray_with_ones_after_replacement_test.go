package sliding_window

import "testing"

func TestLongestSubarrayWithOnesAfterReplacement(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		k    int
		want int
	}{
		"base case: K = 2": {
			in:   []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1},
			k:    2,
			want: 6,
		},
		"base case: k = 3": {
			in:   []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1},
			k:    3,
			want: 9,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := LongestSubarrayWithOnesAfterReplacement(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v: %v want", got, tc.want)
			}
		})
	}
}
