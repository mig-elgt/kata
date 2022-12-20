package sliding_window

func LongestSubstringWithDistinctCharacters(str string) int {
	var maxLength, startWindow int
	subs := map[rune]int{}
	for i := 0; i < len(str); i++ {
		if _, ok := subs[rune(str[i])]; !ok {
			subs[rune(str[i])] = 1
		} else {
			for startWindow < i {
				delete(subs, rune(str[startWindow]))
				startWindow++
			}
			subs[rune(str[i])] = 1
		}
		if maxLength < len(subs) {
			maxLength = len(subs)
		}
	}
	return maxLength
}
