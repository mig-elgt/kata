package reversallinkedlist

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func (this *ListNode) String() string {
	aux := this
	var str string
	for aux != nil {
		str += fmt.Sprintf("%v ", aux.Value)
		aux = aux.Next
	}
	return str
}

func ReverseLinkedList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
