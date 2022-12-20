package bfs

func GetZigzagTraversal(root *TreeNode) [][]int {
	result := [][]int{}
	queue := Queue{}
	queue.Push(root)
	leftToRight := true
	for queue.Len() > 0 {
		levelSize := queue.Len()
		currentLevel := []int{}
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Pop().(*TreeNode)
			if !leftToRight {
				currentLevel = append(currentLevel[:0], append([]int{currentNode.Val}, currentLevel[0:]...)...)
			} else {

				currentLevel = append(currentLevel, currentNode.Val)
			}
			if currentNode.Left != nil {
				queue.Push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.Push(currentNode.Right)
			}
		}
		result = append(result, currentLevel)
		leftToRight = !leftToRight
	}
	return result
}
