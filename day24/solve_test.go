package day24_test

import (
	"testing"

	"github.com/emilioziniades/adventofcode2021/day24"
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

func TestFindAllValidModelNumbers(t *testing.T) {
	changingValues := day24.GetChangingValues("input.txt")
	validModelNumbers := day24.FindAllValidModelNumbers("input.txt")
	want := 0
	for _, modelNumber := range validModelNumbers {
		got := day24.StepAll(modelNumber, changingValues)
		if got != want {
			t.Errorf("TestFindAllValidModelNumbers: got %v, wanted %v, for model number %v", got, want, modelNumber)
		}
	}
}

func TestFindMaxModelNumber(t *testing.T) {
	name := "input.txt"
	got := day24.FindMaxValidModelNumber(name)
	want := 93997999296912
	if got != want {
		t.Errorf("TestFindMaxModelNumber: got %d, wanted %d, for %s", got, want, name)
	}
}

func TestFindMinModelNumber(t *testing.T) {
	name := "input.txt"
	got := day24.FindMinValidModelNumber(name)
	want := 81111379141811
	if got != want {
		t.Errorf("TestFindMinModelNumber: got %d, wanted %d, for %s", got, want, name)
	}

}
