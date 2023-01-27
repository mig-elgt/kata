package grind75

import "testing"

func TestMinKnightMoves(t *testing.T) {
	testCases := []struct {
		x, y, want int
	}{
		{2, 1, 1},
		{5, 5, 4},
	}
	for _, tc := range testCases {
		t.Run("base cases: ", func(t *testing.T) {
			if got := minKnightMoves(tc.x, tc.y); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
