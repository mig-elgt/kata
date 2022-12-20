package two_pointers

import "testing"

func TestTripletsWithSmallerSum(t *testing.T) {
	testCases := map[string]struct {
		in     []int
		target int
		want   int
	}{
		"base case 1": {
			in:     []int{-1, 0, 2, 3},
			target: 3,
			want:   2,
		},
		"base case 2": {
			in:     []int{-1, 4, 2, 1, 3},
			target: 5,
			want:   4,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := TripletsWithSmallerSum(tc.in, tc.target); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
