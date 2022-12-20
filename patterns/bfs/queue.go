package bfs

type Queue []*TreeNode

func (q Queue) Len() int           { return len(q) }
func (q Queue) Less(i, j int) bool { return q[i].Val < q[j].Val }
func (q Queue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(*TreeNode))
}

func (q *Queue) Pop() interface{} {
	old := *q
	x := old[0]
	*q = old[1:]
	return x
}
