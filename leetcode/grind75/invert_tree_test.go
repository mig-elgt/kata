package grind75

import "testing"

func TestInvertTree(t *testing.T) {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	got := invertTree(root)
	wantTree := "9764321"

	if res := got.String(); res != wantTree {
		t.Fatalf("got %v; want %v", res, wantTree)
	}
}
