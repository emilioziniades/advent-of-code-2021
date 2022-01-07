package day7_test

import (
	"testing"

	"github.com/emilioziniades/adventofcode2021/day7"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.FetchData()
}

func TestMinFuel(t *testing.T) {
	testMinFuel(t, "day7-example.txt", 37)
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
