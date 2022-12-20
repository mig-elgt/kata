package topk

import "testing"

func TestKthLargestNumberInStream(t *testing.T) {
	stream := NewNumberStream([]int{3, 1, 5, 12, 2, 11}, 4)

	if got := stream.Add(6); got != 5 {
		t.Fatalf("got %v, want 5", got)
	}

	if got := stream.Add(13); got != 6 {
		t.Fatalf("got %v, want 6", got)
	}

	if got := stream.Add(4); got != 6 {
		t.Fatalf("got %v, want 6", got)
	}
}
