package util

import (
	"fmt"
	"log"
	"math"
)

func Reverse[T any](t []T) {
	for i := len(t)/2 - 1; i >= 0; i-- {
		opp := len(t) - 1 - i
		t[i], t[opp] = t[opp], t[i]
	}
}

func Map[T any](t []T, f func(T) T) []T {
	res := make([]T, 0)
	for _, e := range t {
		res = append(res, f(e))
	}
	return res
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func Has[T comparable](slice []T, item T) bool {
	for _, e := range slice {
		if e == item {
			return true
		}
	}
	return false
}

type Vector interface {
	ToSlice() []int
}

func ManhattanDistance(a, b Vector) (res int) {
	aa := a.ToSlice()
	bb := b.ToSlice()
	if len(aa) != len(bb) {
		log.Fatal("vectors not equal dimensions")
	}
	for i := 0; i < len(aa); i++ {
		res += int(math.Abs(float64(aa[i] - bb[i])))
	}
	return
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	newM := make(map[K]V)
	for k, v := range m {
		newM[k] = v
	}
	return newM
}
