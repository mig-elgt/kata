package linked_list

import "testing"

func Test_LRUCache(t *testing.T) {
	lru := ConstructorLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	value := lru.Get(1)
	if value != 1 {
		t.Fatalf("got %v; want 1 as value", value)
	}
	lru.Put(3, 3)
	// value = lru.Get(2)
	// if value != -1 {
	// 	t.Fatalf("got %v; want -1 as value", value)
	// }
}
