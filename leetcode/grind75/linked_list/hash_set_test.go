package linked_list

import "testing"

func TestHashSet(t *testing.T) {
	hs := NewMyHashSet()
	hs.Add(1)
	hs.Add(2)
	if got := hs.Contains(1); got != true {
		t.Fatalf("got %v; want true", got)
	}
	if got := hs.Contains(3); got != false {
		t.Fatalf("got %v; want false", got)
	}
	hs.Add(2)
	if got := hs.Contains(2); got != true {
		t.Fatalf("got %v; want true", got)
	}
	hs.Remove(2)
	if got := hs.Contains(2); got != false {
		t.Fatalf("got %v; want false", got)
	}
}
