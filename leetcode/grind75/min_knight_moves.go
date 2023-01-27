package grind75

// 1197. Minimum Knight Moves

// In an infinite chess board with coordinates from -infinity to +infinity, you have a knight at square [0, 0].

// A knight has 8 possible moves it can make, as illustrated below. Each move is two squares in a cardinal direction, then one square in an orthogonal direction.

// Return the minimum number of steps needed to move the knight to the square [x, y]. It is guaranteed the answer exists.

// Example 1:
// Input: x = 2, y = 1
// Output: 1
// Explanation: [0, 0] → [2, 1]

// Example 2:
// Input: x = 5, y = 5
// Output: 4
// Explanation: [0, 0] → [2, 1] → [4, 2] → [3, 4] → [5, 5]

// Constraints:
// -300 <= x, y <= 300
// 0 <= |x| + |y| <= 300

import (
	"container/list"
)

func minKnightMoves(x int, y int) int {
	var steps int
	visited := [][]bool{}
	for i := 0; i < 607; i++ {
		row := make([]bool, 607)
		visited = append(visited, row)
	}
	queue := list.New()
	queue.PushBack([]int{0, 0})
	directions := [][]int{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			item := queue.Front()
			curr := item.Value.([]int)
			if curr[0] == x && curr[1] == y {
				return steps
			}
			for _, dir := range directions {
				next := []int{curr[0] + dir[0], curr[1] + dir[1]}
				if !visited[next[0]+302][next[1]+302] {
					visited[next[0]+302][next[1]+302] = true
					queue.PushBack(next)
				}
			}
			queue.Remove(item)
		}
		steps++
	}
	return steps
}
