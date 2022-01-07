package day5_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day5"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/5/input", "day5-input.txt")
	if err != nil {
		log.Fatalf("MapVents: Data: %s", err)
	}
}
func TestMapVents(t *testing.T) {
	testMapVents(t, "day5-example.txt", 5)
	testMapVents(t, "day5-input.txt", 8622)
}

func testMapVents(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("Map Vents: Parsing Input: %s", err)
	}
	got := day5.MapVents(in, false)
	if got != want {
		t.Fatalf("got %d, wanted %d for %s\n", got, want, file)
	}
	fmt.Printf("got %d, wanted %d for %s\n", got, want, file)
}

func TestMapVentsDiag(t *testing.T) {
	testMapVentsDiag(t, "day5-example.txt", 12)
	testMapVentsDiag(t, "day5-input.txt", 22037)
}

func testMapVentsDiag(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("Map Vents Diagonal: %s\n", err)
	}
	got := day5.MapVents(in, true)
	if got != want {
		t.Fatalf("got %d, wanted %d, for %s\n", got, want, file)
	}
	fmt.Printf("got %d, wanted %d, for %s\n", got, want, file)
}
