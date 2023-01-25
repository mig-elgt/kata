package grind75

// 1730. Shortest Path to Get Food
// You are starving and you want to eat food as quickly as possible. You want to find the shortest path to arrive at any food cell.

// You are given an m x n character matrix, grid, of these different types of cells:

// '*' is your location. There is exactly one '*' cell.
// '#' is a food cell. There may be multiple food cells.
// 'O' is free space, and you can travel through these cells.
// 'X' is an obstacle, and you cannot travel through these cells.
// You can travel to any adjacent cell north, east, south, or west of your current location if there is not an obstacle.

// Return the length of the shortest path for you to reach any food cell. If there is no path for you to reach food, return -1.

// Example 1
// Input: grid = [["X","X","X","X","X","X"],["X","*","O","O","O","X"],["X","O","O","#","O","X"],["X","X","X","X","X","X"]]
// Output: 3
// Explanation: It takes 3 steps to reach the food.

// Example 2
// Input: grid = [["X","X","X","X","X"],["X","*","X","O","X"],["X","O","X","#","X"],["X","X","X","X","X"]]
// Output: -1
// Explanation: It is not possible to reach the food.

import (
	"container/list"
)

func getFood(grid [][]byte) int {
	queue := list.New()
out:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '*' {
				queue.PushBack([]int{i, j})
				break out
			}
		}
	}
	var distance int
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			item := queue.Front()
			cell := item.Value.([]int)
			i, j := cell[0], cell[1]
			directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
			for _, dir := range directions {
				ii := i + dir[0]
				jj := j + dir[1]
				if ii >= 0 && ii < len(grid) && jj < len(grid[0]) && jj >= 0 {
					if grid[ii][jj] == '#' {
						return distance + 1
					}
					if grid[ii][jj] == 'O' {
						queue.PushBack([]int{ii, jj})
					}
				}
			}
			grid[i][j] = 'X'
			queue.Remove(item)
		}
		distance++
	}
	return -1
}
