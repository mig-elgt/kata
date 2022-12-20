package two_pointers

import (
	"reflect"
	"testing"
)

func TestTripletSumToZero(t *testing.T) {
	testCases := map[string]struct {
		in   []int
		want [][]int
	}{
		"base case 1": {
			in: []int{-3, 0, 1, 2, -1, 1, -2},
			want: [][]int{
				{-3, 1, 2},
				{-2, 0, 2},
				{-2, 1, 1},
				{-1, 0, 1},
			},
		},
		"base case 2": {
			in: []int{-5, 2, -1, -2, 3},
			want: [][]int{
				{-5, 2, 3},
				{-2, -1, 3},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := TripletSumToZero(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
