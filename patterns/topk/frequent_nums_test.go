package topk

import (
	"reflect"
	"testing"
)

func TestFindTopKFrequentNumbers(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 3, 5, 12, 11, 12, 11}, 2, []int{11, 12}},
		{[]int{5, 12, 11, 3, 11}, 2, []int{12, 11}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindTopKFrequentNumbers(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
