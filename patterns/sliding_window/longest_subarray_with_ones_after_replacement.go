package sliding_window

func LongestSubarrayWithOnesAfterReplacement(arr []int, k int) int {
	var windowStart, maxLength, maxOnesCount int
	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		if arr[windowEnd] == 1 {
			maxOnesCount++
		}
		if windowEnd-windowStart+1-maxOnesCount > k {
			if arr[windowStart] == 1 {
				maxOnesCount--
			}
			windowStart++
		}
		if maxLength < windowEnd-windowStart+1 {
			maxLength = windowEnd - windowStart + 1
		}
	}
	return maxLength
}
