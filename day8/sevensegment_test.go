package day8_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day8"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/8/input", "day8-input.txt")
	if err != nil {
		log.Fatalf("fetch.Data: %s", err)
	}
}

func TestUniqueDigits(t *testing.T) {
	testUniqueDigits(t, "day8-example.txt", 26)
	testUniqueDigits(t, "day8-input.txt", 381)
}
func testUniqueDigits(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("testUniqueDigits: FileToStringSlice: %s", err)
	}

	got := day8.UniqueDigits(in)
	if got != want {
		t.Fatalf("got: %d, want %d for %s", got, want, file)
	}
	t.Logf("got: %d, want %d for %s", got, want, file)
}

func TestFindDigits(t *testing.T) {
	testFindDigits(t, "day8-simple.txt", 5353)
	testFindDigits(t, "day8-example.txt", 61229)
	testFindDigits(t, "day8-input.txt", 1023686)
}

func testFindDigits(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("testFindDigits: FileToStringSlice: %s", err)
	}

	got := day8.FindDigits(in)
	if got != want {
		t.Fatalf("got %d, want %d for %s", got, want, file)
	}
	t.Logf("got: %d, want %d for %s", got, want, file)
}
