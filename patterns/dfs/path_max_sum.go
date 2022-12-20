package dfs

import "math"

func GetPathWithMaxSum(root *TreeNode) int {
	var sum int
	pathWithSumRecursive(root, &sum)
	return sum
}

func pathWithSumRecursive(currentNode *TreeNode, sumMax *int) int {
	if currentNode == nil {
		return 0
	}
	leftSum := pathWithSumRecursive(currentNode.Left, sumMax)
	rigthSum := pathWithSumRecursive(currentNode.Right, sumMax)
	if leftSum != 0 && rigthSum != 0 {
		totalSum := leftSum + rigthSum + currentNode.Val
		if *sumMax < totalSum {
			*sumMax = totalSum
		}
	}
	return int(math.Max(float64(leftSum), float64(rigthSum))) + currentNode.Val
}
