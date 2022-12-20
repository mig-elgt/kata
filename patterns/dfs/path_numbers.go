package dfs

func GetSumPathNumbers(root *TreeNode) int {
	return findSumRecursive(root, 0)
}

func findSumRecursive(root *TreeNode, pathSum int) int {
	if root == nil {
		return 0
	}
	pathSum = 10*pathSum + root.Val
	if root.Left == nil && root.Right == nil {
		return pathSum
	}
	return findSumRecursive(root.Left, pathSum) + findSumRecursive(root.Right, pathSum)
}
