package fastslowpointers

func LinkedListCycleLength(head *ListNode) int {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			var length int
			aux := fast.Next
			for aux != fast {
				aux = aux.Next
				length++
			}
			return length + 1
		}
	}
	return 0
}
