package fastslowpointers

import (
	"reflect"
	"testing"
)

func TestRearrangeLinkedList(t *testing.T) {
	head := &ListNode{Value: 2}
	head.Next = &ListNode{Value: 4}
	head.Next.Next = &ListNode{Value: 6}
	head.Next.Next.Next = &ListNode{Value: 8}
	head.Next.Next.Next.Next = &ListNode{Value: 10}
	head.Next.Next.Next.Next.Next = &ListNode{Value: 12}

	head = RearrangeLinkedList(head)

	rearrangeLinkedListValuesExpected := []int{2, 12, 4, 10, 6, 8}
	got := []int{}

	for head != nil {
		got = append(got, head.Value)
		head = head.Next
	}

	if !reflect.DeepEqual(got, rearrangeLinkedListValuesExpected) {
		t.Fatalf("got %v; want %v", got, rearrangeLinkedListValuesExpected)
	}
}
