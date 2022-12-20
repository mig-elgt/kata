package dynamicprogramming

// Problem: Coin Chainge
// You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money. Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1. You may assume that you have an infinite number of each kind of coin.

// Example 1:

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
// Example 2:

// Input: coins = [2], amount = 3
// Output: -1
// Example 3:

// Input: coins = [1], amount = 0
// Output: 0

import "math"

func coinChange(coins []int, amount int) int {
	// return coinChangeRecursive(coins, amount, map[int]int{})
	return coinChainWithBottomUpApproach(coins, amount)
}

// Brute Force solution
func coinChangeRecursive(coins []int, remain int, cache map[int]int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}
	if _, found := cache[remain]; found {
		return cache[remain]
	}
	minCount := math.MaxInt
	for _, coin := range coins {
		count := coinChangeRecursive(coins, remain-coin, cache)
		if count == -1 {
			continue
		}
		if count+1 < minCount {
			minCount = count + 1
		}
	}
	if minCount == math.MaxInt {
		cache[remain] = -1
	} else {
		cache[remain] = minCount
	}
	return cache[remain]
}

func coinChainWithBottomUpApproach(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
