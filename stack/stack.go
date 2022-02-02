package stack

import "log"

type Stack[T any] []T

func (s *Stack[T]) Pop() T {
	old := *s
	n := len(old)
	if n == 0 {
		log.Fatal("Stack is empty")
	}
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func (s *Stack[T]) PopLeft() T {
	old := *s
	n := len(old)
	if n == 0 {
		log.Fatal("Stack is empty")
	}
	x := old[0]
	*s = old[1:n]
	return x
}

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) PushLeft(x T) {
	*s = append([]T{x}, *s...)
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Reset() {
	*s = New[T]()
}
