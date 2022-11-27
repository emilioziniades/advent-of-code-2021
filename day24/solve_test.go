package day24_test

import (
	"fmt"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day24"
	"github.com/k0kubun/pp/v3"
)

var changingValues = [14][3]int{
	//a, b, c
	{10, 2, 1},   // 0
	{14, 13, 1},  // 1
	{14, 13, 1},  // 2
	{-13, 9, 26}, // 3
	{10, 15, 1},  // 4
	{-13, 3, 26}, // 5
	{-7, 6, 26},  // 6
	{11, 5, 1},   // 7
	{10, 16, 1},  // 8
	{13, 1, 1},   // 9
	{-4, 6, 26},  // 10
	{-9, 3, 26},  // 11
	{-13, 7, 26}, // 12
	{-9, 9, 26},  // 13
}

func TestGetChangingValues(t *testing.T) {
	gotVals := day24.GetChangingValues("input.txt")

	for i, want := range changingValues {
		got := gotVals[i]
		if want[0] != got.A || want[1] != got.B || want[2] != got.C {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestGetRestrictedPairs(t *testing.T) {
	changingValues := day24.GetChangingValues("input.txt")
	pairs := day24.GetRestrictedPairs(changingValues)
	for _, pair := range pairs {
		fmt.Printf("restriction: i%d + %d = i%d\n", pair.Miss.Index, pair.Miss.Values.B+pair.Hit.Values.A, pair.Hit.Index)
	}

}

func TestFindAllModelNumbers(t *testing.T) {
	changingValues := day24.GetChangingValues("input.txt")
	pairs := day24.GetRestrictedPairs(changingValues)
	validModelNumbers := day24.FindAllModelNumbers(pairs)
	pp.Println(validModelNumbers)

}
