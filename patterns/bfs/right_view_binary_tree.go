package bfs

func GetRihtViewBinaryTree(root *TreeNode) []int {
	result := []int{}
	queue := Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		levelSize := queue.Len()
		var rightNode *TreeNode
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Pop().(*TreeNode)
			if currentNode.Left != nil {
				queue.Push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.Push(currentNode.Right)
			}
			rightNode = currentNode
		}
		result = append(result, rightNode.Val)
	}
	return result
}
