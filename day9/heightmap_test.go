package day9_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day9"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	fetch.Data("https://adventofcode.com/2021/day/9/input", "9.in")
}
func TestLowPoints(t *testing.T) {
	testLowPoints(t, "9.ex", 15)
	testLowPoints(t, "9.in", 539)
}

func testLowPoints(t *testing.T, file string, want int) {
	in, err := parse.FileToIntGrid(file)
	if err != nil {
		log.Fatal(err)
	}
	got, _ := day9.LowPoints(in)
	if got != want {
		t.Fatalf("got %d, want %d for %s", got, want, file)
	}
	t.Logf("got %d, want %d for %s", got, want, file)
}

func TestBasinCount(t *testing.T) {
	testBasinCount(t, "9.ex", 1134)
	testBasinCount(t, "9.in", 736920)
}

func testBasinCount(t *testing.T, file string, want int) {
	in, err := parse.FileToIntGrid(file)
	if err != nil {
		log.Fatal(err)
	}

	got := day9.CountBasins(in)
	if got != want {
		t.Fatalf("got %d, want %d for %s", got, want, file)
	}
	t.Logf("got %d, want %d for %s", got, want, file)
}
