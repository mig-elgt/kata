package fastslowpointers

import "testing"

func TestStartOfLinkedListCycle(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	if got := StartOfLinkedListCycle(head); got != head.Next.Next {
		t.Fatalf("got %v; want %v", got, head.Next.Next)
	}

}

func TestStartOfLinkedListCycleSameHead(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}

	head.Next.Next.Next.Next.Next.Next = head

	if got := StartOfLinkedListCycle(head); got != head {
		t.Fatalf("got %v; want %v", got, head.Next.Next)
	}

}
