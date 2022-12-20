package cyclicsort

import "testing"

func TestFindMissingNumber(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{4, 0, 3, 1}, 2},
		{[]int{8, 3, 5, 2, 4, 6, 0, 1}, 7},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := FindMissingNumber(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
