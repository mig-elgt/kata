package reversallinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	head := &ListNode{Value: 2}
	head.Next = &ListNode{Value: 4}
	head.Next.Next = &ListNode{Value: 6}
	head.Next.Next.Next = &ListNode{Value: 8}
	head.Next.Next.Next.Next = &ListNode{Value: 10}

	want := "10 8 6 4 2 "
	got := ReverseLinkedList(head).String()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
