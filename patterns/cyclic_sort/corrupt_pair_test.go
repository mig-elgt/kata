package cyclicsort

import (
	"reflect"
	"testing"
)

func TestFindCorrupPair(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{3, 1, 2, 5, 2}, []int{2, 4}},
		{[]int{3, 1, 2, 3, 6, 4}, []int{3, 5}},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindCorruptPair(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
