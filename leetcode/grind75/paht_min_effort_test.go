package grind75

import "testing"

func TestMinimumEffortPath(t *testing.T) {
	testCases := []struct {
		heights [][]int
		want    int
	}{
		{
			[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}},
			2,
		},
	}
	for _, tc := range testCases {
		t.Run("base case: ", func(t *testing.T) {
			if got := minimumEffortPath(tc.heights); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
