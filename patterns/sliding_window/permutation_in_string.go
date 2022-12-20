package sliding_window

func PermutationInString(str, pattern string) bool {
	var windowStart, matched int
	charFrequency := map[rune]int{}
	for _, chr := range pattern {
		charFrequency[rune(chr)] = 1
	}
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		rightChr := str[windowEnd]
		if _, ok := charFrequency[rune(rightChr)]; ok {
			charFrequency[rune(rightChr)]--
			if charFrequency[rune(rightChr)] == 0 {
				matched++
			}
		}
		if matched == len(pattern) {
			return true
		}
		if windowEnd >= len(pattern)-1 {
			leftChr := str[rune(windowStart)]
			if _, ok := charFrequency[rune(leftChr)]; ok {
				if charFrequency[rune(leftChr)] == 0 {
					matched--
				}
				charFrequency[rune(leftChr)]++
			}
			windowStart++
		}
	}
	return false
}
