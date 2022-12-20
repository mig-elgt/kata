package sliding_window

import "math"

func SmallestWindowContainingSubstring(str, pattern string) string {
	var windowStart, matched, smallestWindowStart int
	smallestWindowLen := math.MaxInt
	charFrequency := map[rune]int{}
	for _, chr := range pattern {
		charFrequency[rune(chr)]++
	}
	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		chr := rune(str[windowEnd])
		if _, ok := charFrequency[chr]; ok {
			charFrequency[chr]--
			if charFrequency[chr] >= 0 {
				matched++
			} else if charFrequency[chr] < 0 {
				for i := windowStart; i < windowEnd; i++ {
					leftChr := rune(str[i])
					if _, ok := charFrequency[leftChr]; ok {
						charFrequency[leftChr]++
						if charFrequency[leftChr] != 0 {
							matched--
						}
					}
					windowStart++
					if leftChr == chr {
						break
					}
				}
			}
		}
		for matched == len(pattern) {
			length := windowEnd - windowStart
			if smallestWindowLen > length {
				smallestWindowLen = length
				smallestWindowStart = windowStart
			}
			if _, ok := charFrequency[rune(str[windowStart])]; ok {
				charFrequency[rune(str[windowStart])]++
				matched--
			}
			windowStart++
		}
	}
	if smallestWindowStart < smallestWindowStart+smallestWindowLen+1 {
		return str[smallestWindowStart : smallestWindowStart+smallestWindowLen+1]
	}
	return ""
}
