package reversallinkedlist

func ReverseAlternatingKElementsSubList(head *ListNode, k int) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		var prev *ListNode
		tailOfSubList := current
		for i := 0; current != nil && i < k; i++ {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
		}
		if previous == nil {
			head, previous = prev, tailOfSubList
			previous.Next = current
		} else {
			tailOfSubList.Next = current
			previous.Next = prev
			previous = tailOfSubList
		}
		for i := 0; current != nil && i < k; i++ {
			previous = current
			current = current.Next
		}
	}
	return head
}
