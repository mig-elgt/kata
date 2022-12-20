package slidingwindow

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1}, 1, []int{1}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	}
	for _, tc := range testCases {
		if got := maxSlidingWindow(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
