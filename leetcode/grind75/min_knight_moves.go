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
	"fmt"
)

func minKnightMoves(x int, y int) int {
	queue := list.New()
	queue.PushBack([]int{0, 0, 0})
	seen := map[string]int{}
	directions := [][]int{{2, 0, 0, 1}, {0, 2, 1, 0}, {-2, 0, 0, 1}, {0, -2, 1, 0}}
	for queue.Len() > 0 {
		item := queue.Front()
		knight := item.Value.([]int)
		if knight[0] == x && knight[1] == y {
			return knight[2]
		}
		x, y := knight[0], knight[1]
		p := fmt.Sprintf("%v,%v", x, y)
		if _, visited := seen[p]; !visited {
			for _, d := range directions {
				carX, carY := x+d[0], y+d[1]
				newOrtX1, newOrthY1 := carX+d[2], carY+d[3]
				newOrtX2, newOrthY2 := (carX+d[2])*-1, (carY+d[3])*-1
				queue.PushBack([]int{newOrtX1, newOrthY1, knight[2] + 1})
				queue.PushBack([]int{newOrtX2, newOrthY2, knight[2] + 1})
			}
			seen[p] = 1
		}
		queue.Remove(item)
	}
	return -1
}
