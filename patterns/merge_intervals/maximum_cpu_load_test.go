package mergeintervals

import "testing"

func TestFindMaxCPULoad(t *testing.T) {
	testCase := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}, 7},
		{[][]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}, 15},
		{[][]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}, 8},
	}
	for _, tc := range testCase {
		t.Run("base cases: ", func(t *testing.T) {
			if got := FindMaxCPULoad(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
