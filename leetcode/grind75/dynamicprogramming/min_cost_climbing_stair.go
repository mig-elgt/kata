package dynamicprogramming

import (
	"math"
)

func minCostClimbingStairs(cost []int) int {
	// return tabulation(cost)
	// return mimimumCost(len(cost), cost, map[int]int{})
	return tabulationOptimized(cost)
}

// Bottom-Up Dynamic Programming (Tabulation)
func tabulation(cost []int) int {
	minimumCost := make([]int, len(cost)+1)
	for i := 2; i < len(minimumCost); i++ {
		oneStep := minimumCost[i-1] + cost[i-1]
		twoSteps := minimumCost[i-2] + cost[i-2]
		minimumCost[i] = int(math.Min(float64(oneStep), float64(twoSteps)))
	}
	return minimumCost[len(minimumCost)-1]
}

// Top-Down Dynamic Programming (Recursion + Memoization)
func mimimumCost(i int, cost []int, memo map[int]int) int {
	if i <= 1 {
		return 0
	}
	if _, cached := memo[i]; cached {
		return memo[i]
	}
	memo[i] = int(math.Min(float64(cost[i-1]+mimimumCost(i-1, cost, memo)), float64(cost[i-2]+mimimumCost(i-2, cost, memo))))
	return memo[i]
}

// Bottom-Up Dynamic Programming Memory Optimized (Tabulation)
func tabulationOptimized(cost []int) int {
	var downOne, downTwo int
	for i := 2; i < len(cost)+1; i++ {
		temp := downOne
		downOne = int(math.Min(float64(downOne+cost[i-1]), float64(downTwo+cost[i-2])))
		downTwo = temp
	}
	return downOne
}
