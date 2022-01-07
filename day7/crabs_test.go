package day7_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day7"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.FetchData("https://adventofcode.com/2021/day/7/input", "day7-input.txt")
	if err != nil {
		log.Fatalf("day7: FetchData: %s", err)
	}
}

func TestMinFuel(t *testing.T) {
	testMinFuel(t, "day7-example.txt", 37)
	testMinFuel(t, "day7-input.txt", 347509)
}

func testMinFuel(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("testMinFuel: FileToStringSlice: %s", err)
	}

	got := day7.MinCost(parse.CommaSeparatedNumbers(in))
	if got != want {
		t.Fatalf("got %d, wanted %d for %s", got, want, file)
	}
	t.Logf("got %d, wanted %d for %s", got, want, file)

}
