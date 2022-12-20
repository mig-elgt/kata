package dynamicprogramming

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	memo := [][]int{}
	for i := 0; i < len(nums)+1; i++ {
		n := make([]int, sum+1)
		for j := 0; j < len(n); j++ {
			n[j] = -1
		}
		memo = append(memo, n)
	}
	return canPartitionRecursion(nums, 0, sum/2, memo)
}

func canPartitionRecursion(nums []int, index, sum int, memo [][]int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || index == len(nums) {
		return false
	}
	if memo[index][sum] != -1 {
		if memo[index][sum] == 1 {
			return true
		} else {
			return false
		}
	}
	res := canPartitionRecursion(nums, index+1, sum-nums[index], memo) || canPartitionRecursion(nums, index+1, sum, memo)
	if res == true {
		memo[index][sum] = 1
	} else {
		memo[index][sum] = 0
	}
	return res
}
