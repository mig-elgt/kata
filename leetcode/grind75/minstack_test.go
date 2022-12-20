package grind75

import "testing"

func TestMinStack(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	if got := minStack.GetMin(); got != 0 {
		t.Fatalf("got %v; want 0", got)
	}
	minStack.Pop()
	if got := minStack.GetMin(); got != 0 {
		t.Fatalf("got %v; want 0", got)
	}
	minStack.Pop()
	if got := minStack.GetMin(); got != 0 {
		t.Fatalf("got %v; want 0", got)
	}
	minStack.Pop()
	if got := minStack.GetMin(); got != 2 {
		t.Fatalf("got %v; want 2", got)
	}
}
