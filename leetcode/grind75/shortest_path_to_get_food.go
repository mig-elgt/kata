package grind75

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
				queue.PushBack([]int{i, j})
				break out
			}
		}
	}
	var distance int
	for queue.Len() > 0 {
		size := queue.Len()
		fmt.Println(size)
		for i := 0; i < size; i++ {
			item := queue.Front()
			cell := item.Value.([]int)
			i, j := cell[0], cell[1]
			fmt.Println(cell)
			directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
			for _, dir := range directions {
				if i+dir[0] < len(grid) && j+dir[1] < len(grid[0]) && i-dir[0] >= 0 && j-dir[1] >= 0 {
					ii := i + dir[0]
					jj := j + dir[1]
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
