package sliding_window

func StringAnagrams(str, pattern string) []int {
	anagramsPositions := []int{}
	charFrequency := map[rune]int{}
	for _, chr := range pattern {
		charFrequency[rune(chr)] = 1
	}
	var windowStart, matched int
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		chr := rune(str[windowEnd])
		if _, ok := charFrequency[chr]; ok {
			charFrequency[chr]--
			if charFrequency[chr] == 0 {
				matched++
			}
		}
		if matched == len(pattern) {
			anagramsPositions = append(anagramsPositions, windowStart)
		}
		if windowEnd-windowStart >= len(pattern)-1 {
			chr := rune(str[windowStart])
			if _, ok := charFrequency[chr]; ok {
				if charFrequency[chr] == 0 {
					matched--
				}
				charFrequency[chr]++
			}
			windowStart++
		}
	}
	return anagramsPositions
}
