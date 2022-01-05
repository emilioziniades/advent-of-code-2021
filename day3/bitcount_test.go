package day3_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day3"
	"github.com/emilioziniades/adventofcode2021/fetch"
)

func init() {
	err := fetch.FetchData("https://adventofcode.com/2021/day/3/input", "day3-input.txt")
	if err != nil {
		log.Fatalf("Fetching data: %s", err)
	}
}
func TestGammaDeltaProd(t *testing.T) {
	testGammaDeltaProd(t, "day3-example.txt", 198)
	testGammaDeltaProd(t, "day3-input.txt", 2261546)
}

func testGammaDeltaProd(t *testing.T, file string, want int) {
	in, err := fetch.ParseInputString(file)
	if err != nil {
		log.Fatalf("GammaDeltaProd: file read error: %s", err)
	}
	prod := day3.GammaDeltaProd(in)
	if prod != want {
		t.Fatalf("GammaDeltaProd: wanted %d, got %d\n", want, prod)
	}
	fmt.Printf("GammaDeltaProd: got %d, wanted %d for %s\n", prod, want, file)
}

func TestOxygenCarbonDioxide(t *testing.T) {
	testOxygenCarbonDioxide(t, "day3-example.txt", 230)
	testOxygenCarbonDioxide(t, "day3-input.txt", 6775520)
}

func testOxygenCarbonDioxide(t *testing.T, file string, want int) {
	in, err := fetch.ParseInputString(file)
	if err != nil {
		log.Fatalf("Oxygen / C02 : file read error: %s", err)
	}
	prod := day3.OxygenCarbonDioxideRating(in)
	if prod != want {
		t.Fatalf("OxygenCarbonDioxideRating: wanted %d, got %d\n", want, prod)
	}
	fmt.Printf("OxygenCarbonDioxideRating: got %d, wanted %d for %s\n", prod, want, file)
}
