package grind75

import "testing"

func TestLinkedList_Reverse(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	reverse := reverseList(head)
	if got, want := reverse.String(), "54321"; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
