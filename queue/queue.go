package queue

type Queue[T any] []T

func (q *Queue[T]) Enqueue(x T) {
	*q = append([]T{x}, *q...)
}

func (q *Queue[T]) Dequeue() T {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

func New[T any]() Queue[T] {
	return Queue[T]{}
}
