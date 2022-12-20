package reversallinkedlist

import (
	"testing"
)

func TestReverseEveryKElementsSunList(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Value: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Value: 8}

	want := "3 2 1 6 5 4 8 7 "
	if got := ReverseEveryKElementsSunList(head, 3).String(); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
