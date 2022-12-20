package kmerge

import (
	"fmt"
)

type List struct {
	Head *ListNode
}

func (list *List) Append(node *ListNode) {
	if list.Head == nil {
		list.Head = node
		return
	}
	tail := list.Head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func (this *ListNode) String() string {
	aux := this
	var str string
	for aux != nil {
		str += fmt.Sprintf("%v, ", aux.Value)
		aux = aux.Next
	}
	return str
}

func MergeKSortedLists(lists []*ListNode) *ListNode {
	panic("not impl")
	// minHeap := heap.NewHeap(func(a, b *ListNode) bool { return a.Value > b.Value })
	// for _, list := range lists {
	// 	if list != nil {
	// 		minHeap.Push(list)
	// 	}
	// }
	// sortedList := List{}
	// for minHeap.Size() > 0 {
	// 	node, _ := minHeap.Pop()
	// 	sortedList.Append(&ListNode{Value: node.Value})
	// 	if node.Next != nil {
	// 		minHeap.Push(node.Next)
	// 	}
	// }
	// return sortedList.Head
}
