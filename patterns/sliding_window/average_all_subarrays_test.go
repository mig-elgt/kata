package sliding_window

import (
	"reflect"
	"testing"
)

func TestAvgAllSubArrys(t *testing.T) {
	testCases := map[string]struct {
		k    int
		in   []int
		want []float64
	}{
		"base case": {
			k:    5,
			in:   []int{1, 3, 2, 6, -1, 4, 1, 8, 2},
			want: []float64{2.2, 2.8, 2.4, 3.6, 2.8},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := AvgAllSubArrays(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
