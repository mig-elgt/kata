package grind75

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  int
	}{
		{5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run("base case: ", func(t *testing.T) {
			if got := countComponents(tt.n, tt.edges); got != tt.want {
				t.Fatalf("got %v; want %v", got, tt.want)
			}
		})
	}
}
