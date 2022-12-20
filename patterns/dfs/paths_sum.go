package dfs

type ResultPaths struct {
	Paths [][]int
}

func GetAllPathsSum(root *TreeNode, sum int) [][]int {
	res := ResultPaths{
		Paths: [][]int{},
	}
	findPathsRecursive(root, sum, &res, []int{root.Val})
	return res.Paths
}

func findPathsRecursive(root *TreeNode, sum int, allPaths *ResultPaths, currentPath []int) {
	if root.Left == nil && root.Right == nil && root.Val == sum {
		allPaths.Paths = append(allPaths.Paths, currentPath)
		return
	}
	if root.Left != nil {
		findPathsRecursive(root.Left, sum-root.Val, allPaths, append(currentPath, root.Left.Val))
	}
	if root.Right != nil {
		findPathsRecursive(root.Right, sum-root.Val, allPaths, append(currentPath, root.Right.Val))
	}
}
