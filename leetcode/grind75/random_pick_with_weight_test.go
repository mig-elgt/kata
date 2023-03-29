package grind75

import "testing"

func TestRandomPickIndex(t *testing.T) {
	sol := NewRandomSolution([]int{2, 4, 8, 1})

	if got := sol.PickIndex(); got != 1 {
		t.Fatalf("got %v; want %v", got, 1)
	}
}
