package cyclicsort

import "testing"

func TestFindDuplicateNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 4, 4, 3, 2}, 4},
		{[]int{2, 1, 3, 3, 5, 4}, 3},
		{[]int{2, 4, 1, 4, 4}, 4},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindDuplicateNumber(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
