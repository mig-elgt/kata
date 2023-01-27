package grind75

// 323. Number of Connected Components in an Undirected Graph
// You have a graph of n nodes. You are given an integer n and an array edges where edges[i] = [ai, bi] indicates that there is an edge between ai and bi in the graph.

// Return the number of connected components in the graph.

// Example 1
// Input: n = 5, edges = [[0,1],[1,2],[3,4]]

// Output: 2

// Example 2
// Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
// Output: 1

// Constraints:

// 1 <= n <= 2000
// 1 <= edges.length <= 5000
// edges[i].length == 2
// 0 <= ai <= bi < n
// ai != bi
// There are no repeated edges.

func countComponents(n int, edges [][]int) int {
	graph := map[int][]int{}
	seen := map[int]int{}
	for i := 0; i < n; i++ {
		graph[i] = []int{}
		seen[i] = 0
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	var graphs int
	for node := 0; node < n; node++ {
		if seen[node] == 0 {
			countRec(node, graph, seen)
			graphs++
		}
	}
	return graphs
}

func countRec(node int, graph map[int][]int, seen map[int]int) {
	if seen[node] == 1 {
		return
	}
	seen[node] = 1
	for _, n := range graph[node] {
		countRec(n, graph, seen)
	}
}
