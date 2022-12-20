package bfs

func CalculateLevelAverage(root *TreeNode) []float64 {
	result := []float64{}
	queue := Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		levelSize := queue.Len()
		levelSum := 0
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Pop().(*TreeNode)
			levelSum += currentNode.Val
			if currentNode.Left != nil {
				queue.Push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.Push(currentNode.Right)
			}
		}
		result = append(result, float64(levelSum)/float64(levelSize))
	}
	return result
}
