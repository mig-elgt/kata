package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func GetLevelOrderTraversal(root *TreeNode) [][]int {
	levels := [][]int{}
	queue := Queue{}
	queue.Push(root)
	for queue.Len() != 0 {
		levelSize := queue.Len()
		currentLevel := []int{}
		for i := 0; i < levelSize; i++ {
			currentNode := queue.Pop().(*TreeNode)
			currentLevel = append(currentLevel, currentNode.Val)
			if currentNode.Left != nil {
				queue.Push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.Push(currentNode.Right)
			}
		}
		levels = append(levels, currentLevel)
	}
	return levels
}
