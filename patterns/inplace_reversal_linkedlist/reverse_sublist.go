package reversallinkedlist

func ReverseSubList(head *ListNode, p, q int) *ListNode {
	current, prevP, nextQ := head.Next, head, head
	var pPnt, qPnt *ListNode
	for current != nil {
		if current.Value == p {
			pPnt = current
		} else if current.Value == q {
			qPnt = current
			nextQ = qPnt.Next
			break
		}
		if pPnt == nil {
			prevP = current
		}
		current = current.Next
	}
	aux := pPnt
	prev := nextQ
	for aux != nextQ {
		next := aux.Next
		aux.Next = prev
		prev = aux
		aux = next
	}
	prevP.Next = qPnt
	return head
}
