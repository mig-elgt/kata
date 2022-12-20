package twoheaps

import "testing"

func TestMedianStream(t *testing.T) {
	medianStream := NewMediaStream()
	medianStream.InsertNum(3)
	medianStream.InsertNum(1)

	if got := medianStream.FindMedian(); got != 2.0 {
		t.Fatalf("got %v; want 2.0", got)
	}

	medianStream.InsertNum(5)

	if got := medianStream.FindMedian(); got != 3.0 {
		t.Fatalf("got %v; want 3.0", got)
	}

	medianStream.InsertNum(4)

	if got := medianStream.FindMedian(); got != 3.5 {
		t.Fatalf("got %v; want 3.0", got)
	}
}
