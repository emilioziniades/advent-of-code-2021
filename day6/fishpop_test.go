package day6_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day6"
	"github.com/emilioziniades/adventofcode2021/fetch"
)

func init() {
	err := fetch.FetchData("https://adventofcode.com/2021/day/6/input", "day6-input.txt")
	if err != nil {
		log.Fatalf("FishPop: FetchData: %s\n", err)
	}
}

func TestFishPop(t *testing.T) {
	testFishPop(t, "day6-example.txt", 5934)
	testFishPop(t, "day6-input.txt", 0)
}

func testFishPop(t *testing.T, file string, want int) {
	in, err := fetch.ParseInputString(file)
	if err != nil {
		t.Fatalf("TestFishPop: ParseInputString: %s\n", err)
	}
	got := day6.SimDays(in, 80)
	if got != want {
		t.Fatalf("FishPop: wanted %d, got %d for %s\n", want, got, file)
	}
	t.Logf("FishPop: wanted %d, got %d for %s\n", want, got, file)

}
