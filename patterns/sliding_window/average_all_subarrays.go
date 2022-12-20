package sliding_window

func AvgAllSubArrays(numbers []int, k int) []float64 {
	result := []float64{}
	var windowSum float64
	var windowStart int
	for windowEnd := 0; windowEnd <= len(numbers)-1; windowEnd++ {
		windowSum += float64(numbers[windowEnd])
		if windowEnd >= k-1 {
			result = append(result, windowSum/float64(k))
			windowSum -= float64(numbers[windowStart])
			windowStart++
		}
	}
	return result
}
