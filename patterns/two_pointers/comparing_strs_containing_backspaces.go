package two_pointers

func ComparingStrsContainingBackspaces(str1, str2 string) bool {
	if applyBackspaces(str1) == applyBackspaces(str2) {
		return true
	}
	return false
}

func applyBackspaces(str string) string {
	strBytes := []rune(str)
	var left int
	for i := 0; i < len(strBytes); i++ {
		if strBytes[i] != rune('#') {
			strBytes[left] = strBytes[i]
			left++
		} else {
			left--
		}
	}
	return string(strBytes[0:left])
}
