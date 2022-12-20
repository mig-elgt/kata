package dfs

func CountPathsForSum(root *TreeNode, target int) int {
	return countPathsRecursive(root, target, []int{})
}

func countPathsRecursive(currentNode *TreeNode, target int, currentPath []int) int {
	if currentNode == nil {
		return 0
	}
	if currentNode.Left == nil && currentNode.Right == nil {
		currentPath = append(currentPath, currentNode.Val)
		var totalPaths, next int
		var sum int
		for i := 0; i < len(currentPath); i++ {
			sum += currentPath[i]
			for sum >= target {
				if sum == target {
					totalPaths++
				}
				sum -= currentPath[next]
				next++
			}
		}
		return totalPaths
	}
	return countPathsRecursive(currentNode.Left, target, append(currentPath, currentNode.Val)) +
		countPathsRecursive(currentNode.Right, target, append(currentPath, currentNode.Val))
}
