package bfs

func ConnectLevelOrderSiblings(root *TreeNode) *TreeNode {
	queue := Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		var previousNode *TreeNode
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Pop().(*TreeNode)
			if previousNode != nil {
				previousNode.Next = currentNode
			}
			previousNode = currentNode
			if currentNode.Left != nil {
				queue.Push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.Push(currentNode.Right)
			}
		}
	}
	return root
}
