package dynamicprogramming

import "testing"

func TestDeleteAndEarn(t *testing.T) {
	testCases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := deleteAndEarn(tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
