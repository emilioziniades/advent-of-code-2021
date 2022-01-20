package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {

	pq := NewPriority[string]()

	items := map[string]int{
		"banana": 5, "apple": 3, "pear": 1,
	}

	for value, priority := range items {
		pq.Enqueue(value, priority)
	}
	if l := pq.Len(); l != 3 {
		t.Fatalf("wanted 3, got %d for length of pq", l)
	}

	item := pq.Dequeue()
	if v, p := item.Value, item.Priority; v != "pear" || p != 1 {
		t.Fatalf("got %s: %d, wanted %s %d for dequeue operation", v, p, "pear", 1)
	}
	pq.Enqueue(item.Value, item.Priority)

	for pq.Len() > 0 {
		curr := pq.Dequeue()
	}
	if l := pq.Len(); l != 0 {
		t.Fatalf("wanted 0, got %d for length of pq", l)
	}
}
