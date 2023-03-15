package grind75

import "testing"

func Test_StackPushPop(t *testing.T) {
	fq := NewFreqStack()

	fq.Push(4)
	fq.Push(0)
	fq.Push(9)
	fq.Push(3)
	fq.Push(4)
	fq.Push(2)

	if got := fq.Pop(); got != 4 {
		t.Fatalf("got %v; want 4", got)
	}

	fq.Push(6)

	if got := fq.Pop(); got != 6 {
		t.Fatalf("got %v; want 6", got)
	}

	fq.Push(1)
	if got := fq.Pop(); got != 1 {
		t.Fatalf("got %v; want 1", got)
	}
	fq.Push(1)
	if got := fq.Pop(); got != 1 {
		t.Fatalf("got %v; want 1", got)
	}
	fq.Push(4)
	if got := fq.Pop(); got != 4 {
		t.Fatalf("got %v; want 4", got)
	}
	if got := fq.Pop(); got != 2 {
		t.Fatalf("got %v; want 2", got)
	}
	if got := fq.Pop(); got != 3 {
		t.Fatalf("got %v; want 3", got)
	}
	if got := fq.Pop(); got != 9 {
		t.Fatalf("got %v; want 9", got)
	}
	if got := fq.Pop(); got != 0 {
		t.Fatalf("got %v; want 0", got)
	}
	if got := fq.Pop(); got != 4 {
		t.Fatalf("got %v; want 4", got)
	}
}
