package reversallinkedlist

func ReverseEveryKElementsSunList(head *ListNode, k int) *ListNode {
	aux := head
	var p *ListNode
	for aux != nil {
		var prev *ListNode
		current := aux
		i := 0
		for current != nil && i < k {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
			i++
		}
		if head == aux {
			head = prev
			p = aux
		} else {
			p.Next = prev
			p = aux
		}
		aux.Next = current
		aux = current
	}
	return head
}
