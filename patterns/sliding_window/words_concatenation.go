package sliding_window

func WordsConcatenation(str string, words []string) []int {
	var windowStart int

	allWords := map[string]int{}

	allWords[words[0]+words[1]] = 1
	allWords[words[1]+words[0]] = 1
	wordsLength := len(words[0] + words[1])

	wordsPositions := []int{}

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		if windowEnd-windowStart == wordsLength-1 {
			if _, ok := allWords[str[windowStart:windowEnd+1]]; ok {
				wordsPositions = append(wordsPositions, windowStart)
			}
			windowStart++
		}
	}

	return wordsPositions
}
