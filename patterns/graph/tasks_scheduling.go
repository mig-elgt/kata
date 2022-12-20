package graph

import "gitlab.com/migel/kata/queue"

func IsSchedulingPossible(tasks int, prerequisites [][]int) bool {
	schedule := map[int][]int{}
	inDegree := map[int]int{}
	// Init schedule tasks graph and in-degree map structures
	for i := 0; i < tasks; i++ {
		schedule[i] = []int{}
		inDegree[i] = 0
	}
	// Build schedule tasks graph
	for _, prerequisite := range prerequisites {
		firstOne, secondOne := prerequisite[0], prerequisite[1]
		firstOnePrerequisites := schedule[firstOne]
		firstOnePrerequisites = append(firstOnePrerequisites, secondOne)
		schedule[firstOne] = firstOnePrerequisites
		inDegree[secondOne]++
	}
	// Find tasks sources
	queue := queue.NewQueue(1)
	for task, value := range inDegree {
		if value == 0 {
			queue.Push(task)
		}
	}
	sortedOrder := []int{}
	// find is schedule tasks graph is scheduing possible
	for queue.Len() > 0 {
		task := queue.Pop()
		sortedOrder = append(sortedOrder, task)
		taskPrerequisites := schedule[task]
		for _, prerequisite := range taskPrerequisites {
			inDegree[prerequisite]--
			if inDegree[prerequisite] == 0 {
				queue.Push(prerequisite)
			}
		}
	}
	return len(sortedOrder) == tasks
}
