package day11_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day11"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/11/input", "11-input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestFlashCount(t *testing.T) {
	var tests = []struct {
		file string
		want int
		days int
	}{
		{"11-simple.txt", 9, 2},
		{"11-example.txt", 204, 10},
		{"11-example.txt", 1656, 100},
		{"11-input.txt", 1729, 100},
	}

	for _, tt := range tests {
		in, err := parse.FileToIntGrid(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day11.FlashCount(in, tt.days)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d for %s", got, tt.want, tt.file)
	}
}
