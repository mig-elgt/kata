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
	"fmt"
)

func getFood(grid [][]byte) int {
	queue := list.New()
out:
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '*' {
				queue.PushBack([]int{i, j, 0})
				break out
			}
		}
	}
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	seen := map[string]int{}
	for queue.Len() > 0 {
		item := queue.Front()
		cell := item.Value.([]int)
		i, j, step := cell[0], cell[1], cell[2]
		if grid[i][j] == '#' {
			return step
		}
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if _, ok := seen[fmt.Sprintf("%v%v", newI, newJ)]; !ok {
				if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] != 'X' {
					queue.PushBack([]int{newI, newJ, step + 1})
				}
			}
		}
		grid[i][j] = 'X'
		queue.Remove(item)
	}
	return -1
}
