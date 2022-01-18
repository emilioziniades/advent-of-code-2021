package day12_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day12"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	fetch.Data("https://adventofcode.com/2021/day/12/input", "12.in")
}

func TestCountPaths(t *testing.T) {
	var tests = []struct {
		file    string
		want    int
		wantTwo int
	}{
		{"12.si", 10, 36},
		{"12.ex", 19, 103},
		{"12.la", 226, 3509},
		{"12.in", 3410, 98796},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}

		got := day12.CountPathsOne(in)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}

		got := day12.CountPathsTwo(in)
		if got != tt.wantTwo {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.wantTwo, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.wantTwo, tt.file)
	}
}
