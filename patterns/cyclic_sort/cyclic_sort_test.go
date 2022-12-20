package cyclicsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{3, 1, 5, 4, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 6, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := Sort(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
