package grind75

import "testing"

func TestMergeKLists(t *testing.T) {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	list3 := &ListNode{Val: 2}
	list3.Next = &ListNode{Val: 6}
	if got, want := mergeKLists([]*ListNode{list1, list2, list3}).String(), "11234456"; got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
