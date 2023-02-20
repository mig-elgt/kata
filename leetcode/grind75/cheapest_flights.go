package grind75

// 787. Cheapest Flights Within K Stops
// There are n cities connected by some number of flights. You are given an array flights where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi to city toi with cost pricei.

// You are also given three integers src, dst, and k, return the cheapest price from src to dst with at most k stops. If there is no such route, return -1.

// Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
// Output: 700
// Explanation:
// The graph is shown above.
// The optimal path with at most 1 stop from city 0 to 3 is marked in red and has cost 100 + 600 = 700.
// Note that the path through cities [0,1,2,3] is cheaper but is invalid because it uses 2 stops.

// Constraints:

// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 104
// There will not be any multiple flights between two cities.
// 0 <= src, dst, k < n
// src != dst

import (
	"container/list"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][][]int{}
	for _, flight := range flights {
		newFligth := []int{flight[1], flight[2]}
		graph[flight[0]] = append(graph[flight[0]], newFligth)
	}
	queue := list.New()
	queue.PushBack([]int{src, 0})
	stops := -1
	cheapest := math.MaxInt
	foundRoute := false
	for queue.Len() > 0 {
		totalFlights := queue.Len()
		for i := 0; i < totalFlights; i++ {
			item := queue.Front()
			currentFligth := item.Value.([]int)
			if stops == k {
				if currentFligth[0] == dst {
					foundRoute = true
					if currentFligth[1] < cheapest {
						cheapest = currentFligth[1]
					}
				}
				queue.Remove(item)
				continue
			}
			for _, child := range graph[currentFligth[0]] {
				queue.PushBack([]int{child[0], currentFligth[1] + child[1]})
			}
			queue.Remove(item)
		}
		stops++
	}
	if foundRoute {
		return cheapest
	}
	return -1
}
