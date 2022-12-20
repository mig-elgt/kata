package fastslowpointers

func StartOfLinkedListCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	distance := LinkedListCycleLength(head)
	for i := 0; i < distance; i++ {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
