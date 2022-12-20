package slidingwindow

func characterReplacement(s string, k int) int {
	lo, hi := 1, len(s)+1
	for lo+1 < hi {
		mid := lo + (hi-lo)/2
		if isValidSubString(s, mid, k) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func isValidSubString(str string, subStrLen, k int) bool {
	freqMap := make([]int, 26)
	var start, maxFrequency int
	for end := 0; end < len(str); end++ {
		freqMap[str[end]-'A']++
		if end+1-start > subStrLen {
			freqMap[str[end]-'A']--
			start++
		}
		if freqMap[str[end]-'A'] > maxFrequency {
			maxFrequency = freqMap[str[end]-'A']
		}
		if subStrLen-maxFrequency <= k {
			return true
		}
	}
	return false
}
