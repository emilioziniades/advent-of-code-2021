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
func TestAStar(t *testing.T) {
	tests := []struct {
		file   string
		want   int
		extend bool
	}{
		{"15.ex", 40, false},
		{"15.in", 472, false},
		{"15.ex", 315, true},
		{"15.in", 2851, true},
	}

	for _, tt := range tests {
		in, err := parse.FileToIntGrid(tt.file)
		if err != nil {
			log.Fatal(err)
		}

		var got int
		if tt.extend {
			got = day15.AStarFivefold(in)
		} else {
			got = day15.AStar(in)
		}
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)

	}
}
