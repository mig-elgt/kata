package dfs

import "testing"

func TestHasPath(t *testing.T) {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	testCases := []struct {
		sum  int
		want bool
	}{
		{23, true},
		{16, false},
		{18, true},
	}
	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := HasPath(root, tc.sum); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
