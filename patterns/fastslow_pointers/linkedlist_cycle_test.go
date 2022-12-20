package fastslowpointers

import "testing"

func TestLinkedListNotCycle(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	if got := LinkedListCycle(head); got != false {
		t.Fatalf("got %v; want false value", got)
	}
}

func TestLinkedListCycle(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	if got := LinkedListCycle(head); got != true {
		t.Fatalf("got %v; want false value", got)
	}
}
