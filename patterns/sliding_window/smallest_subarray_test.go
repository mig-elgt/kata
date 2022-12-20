package sliding_window

import "testing"

func TestSmallestSubArrayWithAGreaterSum(t *testing.T) {
	testCases := map[string]struct {
		S    int
		in   []int
		want int
	}{
		"base case: size 7": {
			S:    7,
			in:   []int{2, 1, 5, 2, 3, 2},
			want: 2,
		},
		"base case: size 8": {
			S:    8,
			in:   []int{3, 4, 1, 1, 6},
			want: 3,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := SmallestSubArrayWithAGreaterSum(tc.in, tc.S); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
