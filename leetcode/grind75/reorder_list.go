package grind75

// 143. Reorder List
// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

// Input: head = [1,2,3,4]
// Output: [1,4,2,3]

func reorderList(head *ListNode) {
	curr := head
	list2 := reverse(mid(head))
	for list2 != nil {
		next2 := list2.Next
		list2.Next = curr.Next
		curr.Next = list2
		curr = curr.Next.Next
		list2 = next2
	}
}

func mid(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr, next := head, head.Next
	for curr != nil {
		curr.Next = prev
		prev = curr
		curr = next
		if curr == nil {
			break
		}
		next = next.Next
	}
	return prev
}
