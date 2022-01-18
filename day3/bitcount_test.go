package day3_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day3"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/3/input", "3.in")
	if err != nil {
		log.Fatalf("Fetching data: %s", err)
	}
}
func TestGammaDeltaProd(t *testing.T) {
	testGammaDeltaProd(t, "3.ex", 198)
	testGammaDeltaProd(t, "3.in", 2261546)
}

func testGammaDeltaProd(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
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
	testOxygenCarbonDioxide(t, "3.ex", 230)
	testOxygenCarbonDioxide(t, "3.in", 6775520)
}

func testOxygenCarbonDioxide(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		log.Fatalf("Oxygen / C02 : file read error: %s", err)
	}
	prod := day3.OxygenCarbonDioxideRating(in)
	if prod != want {
		t.Fatalf("OxygenCarbonDioxideRating: wanted %d, got %d\n", want, prod)
	}
	fmt.Printf("OxygenCarbonDioxideRating: got %d, wanted %d for %s\n", prod, want, file)
}
