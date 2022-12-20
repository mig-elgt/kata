package bfs

func GetLevelOrderSuccessor(root *TreeNode, node int) int {
	queue := Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		currentNode := queue.Pop().(*TreeNode)
		if currentNode.Left != nil {
			queue.Push(currentNode.Left)
		}
		if currentNode.Right != nil {
			queue.Push(currentNode.Right)
		}
		if currentNode.Val == node {
			break
		}
	}
	return queue.Pop().(*TreeNode).Val
}
