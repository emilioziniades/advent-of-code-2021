package day13_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day13"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/13/input", "13-input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestFoldOne(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"13-example.txt", 17},
		{"13-input.txt", 802},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day13.Fold(in, true)
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
		_ = day13.Fold(in, false)
	}
}
