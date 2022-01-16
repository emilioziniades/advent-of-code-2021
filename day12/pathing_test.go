package day12_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day12"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func TestCountPaths(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"12-simple.txt", 10},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}

		got := day12.CountPaths(in)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)
	}
}
