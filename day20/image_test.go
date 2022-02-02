package day20_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day20"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/20/input", "20.in")
	if err != nil {
		log.Fatal(err)
	}
}

func TestEnhance(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{"20.ex", 35},
		{"20.in", 0},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			t.Fatal(err)
		}
		got := day20.Enhance(in, 2)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		} else {
			t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)

		}
	}
}
