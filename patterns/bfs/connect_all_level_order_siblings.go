package bfs

func ConnectAllLevelOrderSiblings(root *TreeNode) *TreeNode {
	queue := Queue{}
	queue.Push(root)
	var previousNode *TreeNode
	for queue.Len() > 0 {
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
	return root
}
