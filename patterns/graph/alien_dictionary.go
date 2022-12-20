package graph

func GetAlienDictionaryOrder(words []string) string {
	panic("not impl")
	// graph := map[rune][]rune{}
	// inDegree := map[rune]int{}
	// // Init graph an in-degree
	// for _, word := range words {
	// 	for _, letter := range word {
	// 		graph[letter] = []rune{}
	// 		inDegree[letter] = 0
	// 	}
	// }
	// // build graph and in-degree
	// for i := 0; i < len(words)-1; i++ {
	// 	wordA := words[i]
	// 	wordB := words[i+1]
	// 	for j := 0; j < len(wordA); j++ {
	// 		parent, child := rune(wordA[j]), rune(wordB[j])
	// 		if parent != child {
	// 			children := graph[parent]
	// 			children = append(children, child)
	// 			graph[parent] = children
	// 			inDegree[child]++
	// 			break
	// 		}
	// 	}
	// }
	// // Find node sources
	// queue := queue.NewQueue(rune('a'))
	// for node, edges := range inDegree {
	// 	if edges == 0 {
	// 		queue.Push(node)
	// 	}
	// }
	// var letters string
	// for queue.Len() > 0 {
	// 	node := queue.Pop()
	// 	letters += string(node)
	// 	children := graph[node]
	// 	for _, child := range children {
	// 		inDegree[child]--
	// 		if inDegree[child] == 0 {
	// 			queue.Push(child)
	// 		}
	// 	}
	// }
	// if len(letters) != len(graph) {
	// 	return ""
	// }
	// return letters
}
