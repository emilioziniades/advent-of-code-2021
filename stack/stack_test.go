package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	s := New[int]()

	s.Push(3)
	s.Push(2)
	s.Push(1)
	fmt.Println(s)

	if x := s.Pop(); x != 1 {
		t.Fatalf("got %d, wanted %d", x, 1)
	}
	fmt.Println(s)

	if x := (Stack[int]{3,2}); !reflect.DeepEqual(x, s) {
		t.Fatalf("got %#v, wanted %#v", s, x)
	}
	fmt.Println(s)

	if x := s.PopLeft(); x != 3 {
		t.Fatalf("got %d, wanted %d", x, 3)
	}
	fmt.Println(s)

	s.PushLeft(4)

	if x := (Stack[int]{4, 2}); !reflect.DeepEqual(x, s) {
		t.Fatalf("got %d, wanted %d", s, x)
	}
	fmt.Println(s)

	s.Reset()
	fmt.Println(s)
}
