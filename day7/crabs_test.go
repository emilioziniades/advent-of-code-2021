package day7_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day7"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/7/input", "day7-input.txt")
	if err != nil {
		log.Fatalf("day7: Data: %s", err)
	}
}

func TestMinFuel(t *testing.T) {
	testMinFuel(t, "day7-example.txt", 37, day7.CostConst)
	testMinFuel(t, "day7-input.txt", 347509, day7.CostConst)
	testMinFuel(t, "day7-example.txt", 168, day7.CostTriangle)
	testMinFuel(t, "day7-input.txt", 98257206, day7.CostTriangle)
}

func testMinFuel(t *testing.T, file string, want int, costfunc func(float64) float64) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("testMinFuel: FileToStringSlice: %s", err)
	}

	got := day7.MinCost(parse.CommaSeparatedNumbers(in), costfunc)
	if got != want {
		t.Fatalf("got %d, wanted %d for %s", got, want, file)
	}
	t.Logf("got %d, wanted %d for %s", got, want, file)

}
