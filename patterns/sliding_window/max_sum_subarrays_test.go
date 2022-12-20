package sliding_window

import "testing"

func TestFindMaxSumSubArrays(t *testing.T) {
	testCases := map[string]struct {
		k    int
		in   []int
		want int
	}{
		"base case: k = 3": {
			k:    3,
			in:   []int{2, 1, 5, 1, 3, 2},
			want: 9,
		},
		"base case: k = 2": {
			k:    2,
			in:   []int{2, 3, 4, 1, 5},
			want: 7,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := FindMaxSumSubArrays(tc.in, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
