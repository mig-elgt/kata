package grind75

//  Number of Islands
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:

// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1

func numIslands(grid [][]byte) int {
	var counter int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				visitIslan(i, j, grid)
				counter++
			}
		}
	}
	return counter
}

func visitIslan(i, j int, grid [][]byte) {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	visitIslan(i, j+1, grid)
	visitIslan(i+1, j, grid)
	visitIslan(i, j-1, grid)
	visitIslan(i-1, j, grid)
}
