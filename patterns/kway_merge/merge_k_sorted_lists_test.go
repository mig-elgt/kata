package kmerge

import (
	"reflect"
	"testing"
)

func TestMergeKSortedLists(t *testing.T) {
	list1 := &ListNode{Value: 2}
	list1.Next = &ListNode{Value: 6}
	list1.Next.Next = &ListNode{Value: 8}

	list2 := &ListNode{Value: 3}
	list2.Next = &ListNode{Value: 6}
	list2.Next.Next = &ListNode{Value: 7}

	list3 := &ListNode{Value: 1}
	list3.Next = &ListNode{Value: 3}
	list3.Next.Next = &ListNode{Value: 4}

	want := "1, 2, 3, 3, 4, 6, 6, 7, 8, "
	if got := MergeKSortedLists([]*ListNode{list1, list2, list3}); !reflect.DeepEqual(got.String(), want) {
		t.Fatalf("got %v; want %v", got.String(), want)
	}
}
