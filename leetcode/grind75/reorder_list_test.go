package grind75

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	want := "1423"

	reorderList(head)
	got := head.String()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%v; want %v", got, want)
	}
}
