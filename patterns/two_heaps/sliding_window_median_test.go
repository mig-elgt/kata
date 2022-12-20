package twoheaps

import (
	"reflect"
	"testing"
)

func TestFindSlidingWindowMedian(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []float64
	}{
		// {[]int{1, 2, -1, 3, 5}, 2, []float64{1.5, 0.5, 1.0, 4.0}},
		{[]int{1, 2, -1, 3, 5}, 3, []float64{1.0, 2.0, 3.0}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindSlidingWindowMedian(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
