package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPath(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return HasPath(root.Left, sum-root.Val) || HasPath(root.Right, sum-root.Val)
}
