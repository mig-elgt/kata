package two_pointers

import "testing"

func TestTripletSumCloseToTarget(t *testing.T) {
	testCases := map[string]struct {
		in     []int
		target int
		want   int
	}{
		"base case 1": {
			in:     []int{-2, 0, 1, 2},
			target: 2,
			want:   1,
		},
		"base case 2": {
			in:     []int{-3, -1, 1, 2},
			target: 1,
			want:   0,
		},
		"base case 3": {
			in:     []int{1, 0, 1, 1},
			target: 100,
			want:   3,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := TripletSumCloseToTarget(tc.in, tc.target); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
