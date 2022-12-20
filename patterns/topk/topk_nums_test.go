package topk

import (
	"reflect"
	"testing"
)

func TestFindKLargestNumbers(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{3, 1, 5, 12, 2, 11}, 3, []int{5, 11, 12}},
		{[]int{5, 12, 11, -1, 12}, 3, []int{11, 12, 12}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindKLargestNumbers(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
