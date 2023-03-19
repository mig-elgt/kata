package grind75

import "testing"

func TestHitCounter(t *testing.T) {
	hitCounter := NewHitCounter()

	hitCounter.Hit(1)
	hitCounter.Hit(2)
	hitCounter.Hit(3)

	if got := hitCounter.GetHits(4); got != 3 {
		t.Fatalf("got %v; want 3 hits", got)
	}

	hitCounter.Hit(300)

	if got := hitCounter.GetHits(300); got != 4 {
		t.Fatalf("got %v; want 4 hits", got)
	}

	if got := hitCounter.GetHits(301); got != 3 {
		t.Fatalf("got %v; want 3 hits", got)
	}
}
