package fastslowpointers

func RearrangeLinkedList(head *ListNode) *ListNode {
	// find the middle of the LinkedList
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow.Next is now pointing the middle one
	headSecondHalf := reverse(slow.Next)
	slow.Next = nil
	headFirstHalf := head
	for headFirstHalf != nil && headSecondHalf != nil {
		revereNext := headSecondHalf.Next
		headSecondHalf.Next = headFirstHalf.Next
		headFirstHalf.Next = headSecondHalf
		headFirstHalf = headFirstHalf.Next.Next
		headSecondHalf = revereNext
	}
	return head
}

func reverse(head *ListNode) *ListNode {
	aux := head
	var prev *ListNode
	for aux != nil {
		next := aux.Next
		aux.Next = prev
		prev = aux
		aux = next
	}
	return prev
}
