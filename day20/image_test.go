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
		file  string
		want  int
		times int
	}{
		{"20.ex", 35, 2},
		{"20.ex", 3351, 50},
		{"20.in", 5354, 2},
		{"20.in", 18269, 50},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			t.Fatal(err)
		}
		got := day20.Enhance(in, tt.times)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		} else {
			t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)

		}
	}
}
