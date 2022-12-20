package cyclicsort

import (
	"reflect"
	"testing"
)

func TestFindAllMissingNumbers(t *testing.T) {
	testCases := []struct {
		in   []int
		want []int
	}{
		{[]int{2, 3, 1, 8, 2, 3, 5, 1}, []int{4, 6, 7}},
		{[]int{2, 4, 1, 2}, []int{3}},
		{[]int{2, 3, 2, 1}, []int{4}},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := FindAllMissingNumbers(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
