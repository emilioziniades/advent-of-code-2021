package day17_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day17"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/17/input", "17.in")
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindMaxY(t *testing.T) {
	var tests = []struct {
		file  string
		maxY  int
		count int
	}{
		{"17.ex", 45, 112},
		{"17.in", 4560, 3344},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		maxY, count := day17.FindMaxY(in[0])
		if maxY != tt.maxY {
			t.Fatalf("maximum Y: got %d, want %d, for %s", maxY, tt.maxY, tt.file)
		}
		t.Logf("maximum Y: got %d, want %d, for %s", maxY, tt.maxY, tt.file)
		if count != tt.count {
			t.Fatalf("count: got %d, want %d, for %s", count, tt.count, tt.file)
		}
		t.Logf("count: got %d, want %d, for %s", count, tt.count, tt.file)
	}
}
