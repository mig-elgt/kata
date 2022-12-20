package dfs

func HasPathSequence(root *TreeNode, seq []int) bool {
	return findPathSeqRecursive(root, seq, 0)
}

func findPathSeqRecursive(currentNode *TreeNode, seq []int, seqIndex int) bool {
	if currentNode == nil {
		return false
	}
	if seqIndex >= len(seq) || currentNode.Val != seq[seqIndex] {
		return false
	}
	if currentNode.Left == nil && currentNode.Right == nil && seqIndex == len(seq)-1 {
		return true
	}
	return findPathSeqRecursive(currentNode.Left, seq, seqIndex+1) ||
		findPathSeqRecursive(currentNode.Right, seq, seqIndex+1)
}
