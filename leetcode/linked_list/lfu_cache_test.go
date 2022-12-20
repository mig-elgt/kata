package linked_list

import "testing"

func TestLFUCacheBaseCase(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	if got := lfu.Get(1); got != 1 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	lfu.Put(3, 3)
	if got := lfu.Get(2); got != -1 {
		t.Fatalf("get(2): got %v; want -1", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Fatalf("get(3): got %v; want 3", got)
	}
	lfu.Put(4, 4)
	if got := lfu.Get(1); got != -1 {
		t.Fatalf("get(1): got %v; want -1", got)
	}
	if got := lfu.Get(3); got != 3 {
		t.Fatalf("get(3): got %v; want 3", got)
	}
	if got := lfu.Get(4); got != 4 {
		t.Fatalf("get(4): got %v; want 4", got)
	}
}

func TestLFUCacheCapacity3(t *testing.T) {
	lfu := Constructor(3)
	lfu.Put(2, 2)
	lfu.Put(1, 1)
	if got := lfu.Get(2); got != 2 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	if got := lfu.Get(1); got != 1 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	if got := lfu.Get(2); got != 2 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	if got := lfu.Get(3); got != -1 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	if got := lfu.Get(2); got != 2 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	if got := lfu.Get(1); got != 1 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
	if got := lfu.Get(4); got != 4 {
		t.Fatalf("get(1): got %v; want 1", got)
	}
}

func TestLFUCacheCapacityZero(t *testing.T) {
	lfu := Constructor(0)
	lfu.Put(0, 0)
	if got := lfu.Get(0); got != -1 {
		t.Fatalf("get(1): got %v; want -1", got)
	}
}

func TestLFUUpdateSameKey(t *testing.T) {
	lfu := Constructor(2)
	if got := lfu.Get(2); got != -1 {
		t.Fatalf("get(2): got %v; want -1", got)
	}
	lfu.Put(2, 6)
	if got := lfu.Get(1); got != -1 {
		t.Fatalf("get(1): got %v; want -1", got)
	}
	lfu.Put(1, 5)
	lfu.Put(1, 2)
	if got := lfu.Get(1); got != 2 {
		t.Fatalf("get(1): got %v; want 2", got)
	}
	if got := lfu.Get(2); got != 6 {
		t.Fatalf("get(1): got %v; want 6", got)
	}
}
