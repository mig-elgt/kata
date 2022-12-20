package reversallinkedlist

import (
	"reflect"
	"testing"
)

func TestReverseSubList(t *testing.T) {
	head := &ListNode{Value: 1}
	head.Next = &ListNode{Value: 2}
	head.Next.Next = &ListNode{Value: 3}
	head.Next.Next.Next = &ListNode{Value: 4}
	head.Next.Next.Next.Next = &ListNode{Value: 5}

	want := "1 4 3 2 5 "
	got := ReverseSubList(head, 2, 4).String()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v; want %v", got, want)
	}
}
