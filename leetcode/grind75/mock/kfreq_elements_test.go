package mock

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFreqNums(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{2, 2, 2, 3, 2, 2, 3, 5}, 2, []int{2, 3}},
	}
	for _, tc := range testCases {
		t.Run("base cases:", func(t *testing.T) {
			got := TopKFreqNums(tc.nums, tc.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
