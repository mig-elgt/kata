package sliding_window

func FruitsIntoBaskets(fruits []rune) int {
	var maxLength, windowStart int
	baskets := map[rune]int{}
	for i := 0; i < len(fruits); i++ {
		baskets[fruits[i]]++
		for len(baskets) > 2 {
			fruit := fruits[windowStart]
			baskets[fruit]--
			if baskets[fruit] == 0 {
				delete(baskets, fruits[windowStart])
			}
			windowStart++
		}
		if maxLength < i-windowStart+1 {
			maxLength = i - windowStart + 1
		}
	}
	return maxLength
}
