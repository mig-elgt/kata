package topk

func SortCharacterByFrequency(str string) string {
	panic("not impl")
	// chrsFrequency := map[rune]int{}
	// for _, chr := range str {
	// 	chrsFrequency[chr]++
	// }
	// heap := heap.NewHeap(func(a, b rune) bool { return chrsFrequency[a] < chrsFrequency[b] })
	// for chr := range chrsFrequency {
	// 	heap.Push(chr)
	// }
	// var result string
	// for heap.Size() > 0 {
	// 	char, _ := heap.Pop()
	// 	for i := 0; i < chrsFrequency[char]; i++ {
	// 		result += string(char)
	// 	}
	// }
	// return result
}
