package fastslowpointers

import "testing"

func TestLinkedListCycleLength(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	if got := LinkedListCycleLength(head); got != 4 {
		t.Fatalf("got %v; want 4 length", got)
	}
}

func TestLinkedListCycleLengthSameHeadt(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head.Next.Next.Next.Next.Next.Next = head

	if got := LinkedListCycleLength(head); got != 6 {
		t.Fatalf("got %v; want 4 length", got)
	}
}
