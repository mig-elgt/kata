package queue

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](a T) *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (gq Queue[T]) Len() int { return len(gq.data) }

func (gq *Queue[T]) Push(x T) {
	gq.data = append(gq.data, x)
}

func (gq *Queue[T]) Pop() T {
	val := gq.data[0]
	gq.data = gq.data[1:]
	return val
}
