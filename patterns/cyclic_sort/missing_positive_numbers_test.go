package cyclicsort

import (
	"reflect"
	"testing"
)

func TestFindFirstKMissingPositiveNumbers(t *testing.T) {
	testCases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{3, -1, 4, 5, 5}, 3, []int{1, 2, 6}},
		{[]int{2, 3, 4}, 3, []int{1, 5, 6}},
		{[]int{-2, -3, 4}, 2, []int{1, 2}},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindFirstKMissingPositiveNumbers(tc.in, tc.k); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
