package graph

func SortTopologicalGraph(vertices int, edges ...[]int) []int {
	panic("not impl")
	// sortedOrder := []int{}
	// graph, inDegree := buildGraphAndInDegree(vertices, edges...)
	// nodeSources := findNodeSources(inDegree)
	// queue := queue.NewQueue(0)
	// for _, source := range nodeSources {
	// 	queue.Push(source)
	// }
	// for queue.Len() > 0 {
	// 	vertice := queue.Pop()
	// 	sortedOrder = append(sortedOrder, vertice)
	// 	children := graph[vertice]
	// 	for _, child := range children {
	// 		inDegree[child]--
	// 		if inDegree[child] == 0 {
	// 			queue.Push(child)
	// 		}
	// 	}
	// }
	// // topological sort is not possible as the graph has a cycle
	// if len(sortedOrder) != vertices {
	// 	return []int{}
	// }
	// return sortedOrder
}

func findNodeSources(inDegree map[int]int) []int {
	sources := []int{}
	for vertice, edges := range inDegree {
		if edges == 0 {
			sources = append(sources, vertice)
		}
	}
	return sources
}

func buildGraphAndInDegree(vertices int, edges ...[]int) (map[int][]int, map[int]int) {
	graph := map[int][]int{}
	inDegree := map[int]int{}
	for i := 0; i < vertices; i++ {
		graph[i] = []int{}
		inDegree[i] = 0
	}
	for _, edge := range edges {
		parent := edge[0]
		children := graph[parent]
		children = append(children, edge[1])
		graph[parent] = children
		inDegree[edge[1]]++
	}
	return graph, inDegree
}
