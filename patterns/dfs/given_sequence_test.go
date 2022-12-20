package dfs

import "testing"

func TestHasPathSequence(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 5}

	testCases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 0, 7}, false},
		{[]int{1, 1, 6}, true},
	}

	for _, tc := range testCases {
		t.Run("base cases", func(t *testing.T) {
			if got := HasPathSequence(root, tc.in); got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
