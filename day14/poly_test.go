package day14_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day14"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/14/input", "14.in")
	if err != nil {
		log.Fatal(err)
	}
}
func TestPolymer(t *testing.T) {
	var tests = []struct {
		file  string
		want  int
		steps int
	}{
		{"14.ex", 1588, 10},
		{"14.in", 2947, 10},
		{"14.ex", 2188189693529, 40},
		{"14.in", 3232426226464, 40},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day14.Polymer(in, tt.steps)
		if got != tt.want {
			t.Fatalf("got %d, want %d, after %d steps, for %s", got, tt.want, tt.steps, tt.file)
		}
		t.Logf("got %d, want %d, after %d steps, for %s", got, tt.want, tt.steps, tt.file)
	}
}
