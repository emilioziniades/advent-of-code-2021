package day9_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day9"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	fetch.Data("https://adventofcode.com/2021/day/9/input", "day9-input.txt")
}
func TestLowPoints(t *testing.T) {
	testLowPoints(t, "day9-example.txt", 15)
	testLowPoints(t, "day9-input.txt", 0)
}

func testLowPoints(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		log.Fatal(err)
	}
	got := day9.LowPoints(in)
	if got != want {
		t.Fatalf("got %d, want %d for %s", got, want, file)
	}
	t.Logf("got %d, want %d for %s", got, want, file)
}
