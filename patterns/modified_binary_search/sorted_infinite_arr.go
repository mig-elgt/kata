package binarysearch

import "math"

type ArrayReader struct {
	arr []int
}

func NewArrayReader(arr []int) ArrayReader {
	return ArrayReader{arr}
}

func (this ArrayReader) Get(index int) int {
	if index >= len(this.arr) {
		return math.MaxInt
	}
	return this.arr[index]
}

func SearchSortedInfiniteArray(reader ArrayReader, key int) int {
	start, end := 0, math.MaxInt
	for start <= end {
		middle := start + (end-start)/2
		num := reader.Get(middle)
		if num == math.MaxInt {
			end = middle - 1
		} else if key > num {
			start = middle + 1
		} else if key < num {
			end = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
