package util

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	i := []int{1, 2, 3, 4}
	g := Map(i, func(i int) int { return i * 5 })
	w := []int{5, 10, 15, 20}
	if !reflect.DeepEqual(g, w) {
		t.Errorf("wanted %v, got %v", w, g)
	}
}
