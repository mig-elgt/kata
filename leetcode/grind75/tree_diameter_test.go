package grind75

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	wantDiameter := 3
	if got := diameterOfBinaryTree(root); got != wantDiameter {
		t.Fatalf("got %v; want %v", got, wantDiameter)
	}

}
