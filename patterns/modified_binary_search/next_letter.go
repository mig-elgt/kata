package binarysearch

func FindNextLetter(arr []rune, key rune) rune {
	start, end := 0, len(arr)-1
	for start <= end {
		middle := start + (end-start)/2
		if key < arr[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	return arr[start%len(arr)]
}
