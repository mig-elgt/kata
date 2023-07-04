package bfs

import "testing"

func TestIsSubPath(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 1}
	root.Right.Left.Left = &TreeNode{Val: 6}
	root.Right.Left.Right = &TreeNode{Val: 8}
	root.Right.Left.Right.Left = &TreeNode{Val: 1}
	root.Right.Left.Right.Right = &TreeNode{Val: 3}

	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 9}
	head.Next.Next = &ListNode{Val: 3}

	if got := isSubPath(head, root); got != false {
		t.Fatalf("got %v; want false", got)
	}
}
