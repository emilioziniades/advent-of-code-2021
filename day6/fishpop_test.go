package day6_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day6"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

var benchIn []string
var benchDays int

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/6/input", "day6-input.txt")
	if err != nil {
		log.Fatalf("FishPop: Data: %s\n", err)
	}

	benchIn, err = parse.FileToStringSlice("day6-input.txt")
	if err != nil {
		log.Fatalf("benchIn: %s", err)
	}

	benchDays = 100
}

func TestFishPop(t *testing.T) {
	testFishPopDict(t, "day6-example.txt", 5934, 80)
	testFishPopDict(t, "day6-input.txt", 383160, 80)
	testFishPopDict(t, "day6-example.txt", 26984457539, 256)
	testFishPopDict(t, "day6-input.txt", 1721148811504, 256)
}

func testFishPop(t *testing.T, file string, want int, days int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("TestFishPop: FileToStringSlice: %s\n", err)
	}
	got := day6.FishPop(in, days)
	if got != want {
		t.Fatalf("FishPop: wanted %d, got %d for %s after %d days\n", want, got, file, days)
	}
	t.Logf("FishPop: wanted %d, got %d for %s after %d days\n", want, got, file, days)

}

func testFishPopDict(t *testing.T, file string, want int, days int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		t.Fatalf("TestFishPop: FileToStringSlice: %s\n", err)
	}
	got := day6.FishPopDict(in, days)
	if got != want {
		t.Fatalf("FishPop: wanted %d, got %d for %s after %d days\n", want, got, file, days)
	}
	t.Logf("FishPop: wanted %d, got %d for %s after %d days\n", want, got, file, days)

}

func BenchmarkFishPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day6.FishPop(benchIn, benchDays)
	}
}

func BenchmarkFishPopDict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = day6.FishPopDict(benchIn, benchDays)
	}
}
