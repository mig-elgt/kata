package grind75

import "testing"

func TestLinkedList_HasCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = head.Next

	want := true
	if got := hasCycle(head); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestLinkedList_HasNotCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	want := false
	if got := hasCycle(head); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
