package day8_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day8"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/8/input", "8.in")
	if err != nil {
		log.Fatalf("fetch.Data: %s", err)
	}
}

func TestUniqueDigits(t *testing.T) {
	testUniqueDigits(t, "8.ex", 26)
	testUniqueDigits(t, "8.in", 381)
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
	testFindDigits(t, "8.si", 5353)
	testFindDigits(t, "8.ex", 61229)
	testFindDigits(t, "8.in", 1023686)
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
