package topk

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	testCases := []struct {
		nums       []int
		key, value int
		want       []int
	}{
		{[]int{5, 6, 7, 8, 9}, 3, 7, []int{6, 8, 7}},
		{[]int{2, 4, 5, 6, 9}, 3, 6, []int{4, 5, 6}},
		{[]int{2, 4, 5, 6, 9}, 3, 10, []int{5, 6, 9}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindClosestElements(tc.nums, tc.key, tc.value); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
