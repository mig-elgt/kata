package grind75

import "testing"

func TestIsValidBSTBaseCase(t *testing.T) {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	want := true
	if got := isValidBST(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestIsValidBSTWithBadNodes(t *testing.T) {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}

	want := false
	if got := isValidBST(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

func TestIsValidBSTWithTwoNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	want := false
	if got := isValidBST(root); got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
