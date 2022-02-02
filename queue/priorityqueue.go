package queue

import "container/heap"

type Item[T any] struct {
	Value    T
	Priority int
	index    int
}

type PriorityQueue[T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func NewPriority[T any]() PriorityQueue[T] {

	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)
	return pq
}

func (pq *PriorityQueue[T]) Enqueue(x T, p int) {
	item := Item[T]{
		Value:    x,
		Priority: p,
	}
	heap.Push(pq, &item)
}
func (pq *PriorityQueue[T]) Dequeue() Item[T] {
	x := heap.Pop(pq).(*Item[T])
	return *x
}
