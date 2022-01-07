package day4_test

import (
	"fmt"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day4"
	"github.com/emilioziniades/adventofcode2021/fetch"
)

func init() {
	fetch.Data("https://adventofcode.com/2021/day/4/input", "day4-input.txt")
}
func TestBingo(t *testing.T) {
	testBingo(t, "day4-example.txt", 4512)
	testBingo(t, "day4-input.txt", 55770)
}
func testBingo(t *testing.T, file string, want int) {
	boards, nums, err := day4.ParseBingo(file)
	if err != nil {
		t.Fatalf("ParseBingo: %s", err)
	}
	score := day4.PlayBingo(nums, boards)
	if score != want {
		t.Fatalf("PlayBingo: wanted %d, got %d for %s\n", want, score, file)
	}
	fmt.Printf("PlayBingo: wanted %d, got %d for %s\n", want, score, file)
}

func TestLoseBingo(t *testing.T) {
	testLoseBingo(t, "day4-example.txt", 1924)
	testLoseBingo(t, "day4-input.txt", 2980)
}
func testLoseBingo(t *testing.T, file string, want int) {
	boards, nums, err := day4.ParseBingo(file)
	if err != nil {
		t.Fatalf("ParseBingo: %s", err)
	}

	score := day4.LoseBingo(nums, boards)
	if score != want {
		t.Fatalf("LoseBingo: wanted %d, got %d for %s\n", want, score, file)
	}
	fmt.Printf("LoseBingo: wanted %d, got %d for %s\n", want, score, file)
}
