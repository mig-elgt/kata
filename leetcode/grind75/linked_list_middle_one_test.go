package grind75

import "testing"

func TestLinkedList_MiddleNodeWithOddNodes(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	if got, want := middleNode(head), head.Next.Next; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestLinkedList_MiddleNodeWithPairNodes(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	if got, want := middleNode(head), head.Next.Next.Next; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
