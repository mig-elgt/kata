package linked_list

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	Head *Node
}

func (list *List) Append(node *Node) {
	if list.Head == nil {
		list.Head = node
		return
	}
	aux := list.Head
	prev := list.Head
	for aux != nil {
		prev = aux
		aux = aux.Next
	}
	prev.Next = node
}

func IsAcyclic(list *List) bool {
	if list == nil {
		return true
	}
	aux := list.Head
	if aux == nil {
		return true
	}
	visited := map[*Node]bool{}
	for aux != nil {
		if _, found := visited[aux]; found {
			return false
		}
		visited[aux] = true
		aux = aux.Next
	}
	return true
}
