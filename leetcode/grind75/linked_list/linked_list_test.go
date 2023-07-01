package linked_list

import "testing"

func TestMyLinkedList(t *testing.T) {
	list := NewMyLinkedList()
	list.AddAtIndex(0, 10)
	list.AddAtIndex(0, 20)
	list.AddAtIndex(1, 30)
	if got := list.Get(0); got != 20 {
		t.Fatalf("got %v; want 20", got)
	}
}
