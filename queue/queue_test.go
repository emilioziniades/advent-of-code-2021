package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()
	fmt.Println(q)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q)

	if !reflect.DeepEqual(q, Queue[int]{3, 2, 1}) {
		t.Fatalf("got %v, wanted %v", q, []int{3, 2, 1})
	}

	curr := q.Dequeue()
	fmt.Println(q)

	if curr != 1 {
		t.Fatalf("got %d, wanted %d", curr, 1)
	}

	qs := New[string]()
	qs.Enqueue("o")
	qs.Enqueue("l")
	qs.Enqueue("l")
	qs.Enqueue("e")
	qs.Enqueue("H")

	fmt.Println(qs)

	if !reflect.DeepEqual(qs, Queue[string]{"H", "e", "l", "l", "o"}) {
		t.Fatalf("got %v, wanted %v", q, []int{3, 2, 1})
	}
	//	q.Enqueue("Haha generics")

}
