package isogram

import "strings"

func IsIsogram(word string) bool {
	if word == "" {
		return true
	}
	w := strings.ToLower(word)
	for i := 0; i < len(w)-1; i++ {
		if w[i] != '-' && w[i] != ' ' {
			for j := i + 1; j < len(w); j++ {
				if w[i] == w[j] {
					return false
				}
			}
		}
	}
	return true
}
