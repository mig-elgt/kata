package greedy

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := canCompleteCircuit(tc.gas, tc.cost); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
