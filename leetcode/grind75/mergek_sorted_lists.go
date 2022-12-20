package grind75

// Problem: Merge k Sorted Lists
// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

// Merge all the linked-lists into one sorted linked-list and return it.

// Example 1:

// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

import "container/heap"

type MinNodeHeap []*ListNode

func (h MinNodeHeap) Len() int           { return len(h) }
func (h MinNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := MinNodeHeap{}
	heap.Init(&minHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}
	var tail, head *ListNode
	for minHeap.Len() > 0 {
		node := heap.Pop(&minHeap).(*ListNode)
		newNode := &ListNode{Val: node.Val}
		if head == nil {
			head = newNode
		} else {
			tail.Next = newNode
		}
		tail = newNode
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
	}
	return head
}
