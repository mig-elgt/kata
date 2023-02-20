package grind75

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	testCases := []struct {
		n           int
		flights     [][]int
		src, dst, k int
		want        int
	}{
		{
			4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1, 700,
		},
		{

			3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1, 200,
		},
		{

			3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0, 500,
		},
	}
	for _, tc := range testCases {
		t.Run("base case", func(t *testing.T) {
			if got := findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
