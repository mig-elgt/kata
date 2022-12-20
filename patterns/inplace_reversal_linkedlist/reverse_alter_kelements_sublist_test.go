package reversallinkedlist

import "testing"

func TestReverseAlternatingKElementsSunList(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Value: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Value: 8}

	want := "2 1 3 4 6 5 7 8 "
	if got := ReverseAlternatingKElementsSubList(head, 2).String(); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}

}
