package grind75

// Reverse Linked List
// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Example 1
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	aux := head
	for aux != nil {
		next = aux.Next
		aux.Next = prev
		prev = aux
		aux = next
	}
	return prev
}
