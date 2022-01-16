package day11

type point struct {
	R, C int
}
type queue []point

func (q *queue) Enqueue(x point) {
	*q = append(queue{x}, *q...)
}

func (q *queue) Dequeue() point {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
