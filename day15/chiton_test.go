package day15_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day15"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/15/input", "15.in")
	if err != nil {
		log.Fatal(err)
	}
}
func TestDijkstra(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{"15.ex", 40},
		{"15.in", 472},
	}

	for _, tt := range tests {
		in, err := parse.FileToIntGrid(tt.file)
		if err != nil {
			log.Fatal(err)
		}

		got := day15.Dijkstra(in)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)

	}
}
