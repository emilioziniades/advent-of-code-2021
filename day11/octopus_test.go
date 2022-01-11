package day11_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day11"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func TestFlashCount(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"11-example.txt", 1656},
	}

	for _, tt := range tests {
		in, err := parse.FileToIntGrid(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day11.FlashCount(in, 10)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d for %s", got, tt.want, tt.file)
	}
}
