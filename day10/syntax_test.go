package day10_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day10"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	fetch.Data("https://adventofcode.com/2021/day/10/input", "10.in")
}

func TestErrorScore(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"10.ex", 26397},
		{"10.in", 341823},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day10.ErrorScore(in)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)
	}

}

func TestCompletionScore(t *testing.T) {
	var tests = []struct {
		file string
		want int
	}{
		{"10.ex", 288957},
		{"10.in", 2801302861},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day10.CompletionScore(in)
		if got != tt.want {
			t.Fatalf("got %d, wanted %d, for %s", got, tt.want, tt.file)
		}
		t.Logf("got %d, wanted %d, for %s", got, tt.want, tt.file)
	}

}
