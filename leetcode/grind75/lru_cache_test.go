package grind75

import "testing"

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	if got := lru.Get(1); got != 1 {
		t.Fatalf("got %v; want 1", got)
	}

	lru.Put(3, 3)
	if got := lru.Get(2); got != -1 {
		t.Fatalf("got %v; want -1", got)
	}

	lru.Put(4, 4)
	if got := lru.Get(1); got != -1 {
		t.Fatalf("got %v; want -1", got)
	}

	if got := lru.Get(3); got != 3 {
		t.Fatalf("got %v; want 3", got)
	}

	if got := lru.Get(4); got != 4 {
		t.Fatalf("got %v; want 4", got)
	}
}
