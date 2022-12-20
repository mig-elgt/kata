package grind75

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{2, 1}},
		{[]int{1}, 1, []int{1}},
	}
	for _, tc := range testCases {
		if got := topKFrequentElements(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %v; want %v", got, tc.want)
		}
	}
}
