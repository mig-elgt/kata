package trees

type Node struct {
	Left  *Node
	Rigth *Node
	Value int
}

type tree struct {
	Head *Node
}

func (t tree) Height() int {
	aux := t.Head
	return t.getHeight(aux)
}

func (t tree) getHeight(node *Node) int {
	if node == nil {
		return 0
	}
	leftCounter := t.getHeight(node.Left)
	rigthCounter := t.getHeight(node.Rigth)
	max := leftCounter
	if rigthCounter > leftCounter {
		max = rigthCounter
	}
	return max + 1
}

func (t *tree) Insert(value int) {
	node := &Node{Value: value}
	if t.Head == nil {
		t.Head = node
		return
	}
	aux := t.Head
	t.insertRec(aux, node)
}

func (t *tree) insertRec(current, node *Node) bool {
	if current == nil {
		return true
	}
	if node.Value <= current.Value {
		if t.insertRec(current.Left, node) {
			current.Left = node
		}
	} else {
		if t.insertRec(current.Rigth, node) {
			current.Rigth = node
		}
	}
	return false
}
