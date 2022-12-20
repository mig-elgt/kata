package sliding_window

func LongestSubWithMaxKDistinct(str string, k int) int {
	var longest, startWindow int
	subs := map[rune]int{}
	for i := 0; i < len(str); i++ {
		subs[rune(str[i])] += 1
		for len(subs) > k {
			if longest < i-1-startWindow {
				longest = (i - 1 - startWindow) + 1
			}
			letter := rune(str[startWindow])
			subs[letter] -= 1
			if subs[letter] == 0 {
				delete(subs, letter)
			}
			startWindow++
		}
	}
	if len(subs) < k {
		return len(str)
	}
	return longest
}
