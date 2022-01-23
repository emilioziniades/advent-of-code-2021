package day16_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day16"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/16/input", "16.in")
	if err != nil {
		log.Fatal(err)
	}
}

func TestBits(t *testing.T) {
	tests := []struct {
		file         string
		versionCount int
		result       int
	}{
		{"16.si", 6, 2021},
		{"16.si2", 9, 1},
		{"16.si3", 14, 3},
		{"16.ex", 16, 15},
		{"16.ex2", 12, 46},
		{"16.la", 23, 46},
		{"16.la2", 31, 54},
		{"16.mo", 14, 3},
		{"16.mo2", 8, 54},
		{"16.mo3", 15, 7},
		{"16.mo4", 11, 9},
		{"16.mo5", 13, 1},
		{"16.mo6", 19, 0},
		{"16.mo7", 16, 0},
		{"16.mo8", 20, 1},
		{"16.in", 847, 333794664059},
	}

	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		versionCount, result := day16.Bits(in)
		if versionCount != tt.versionCount {
			t.Fatalf("version count: got %d, wanted %d, for %s", versionCount, tt.versionCount, tt.file)
		} else {
			t.Logf("version count: got %d, wanted %d, for %s", versionCount, tt.versionCount, tt.file)
		}
		if result != tt.result {
			t.Fatalf("result: got %d, wanted %d, for %s", result, tt.result, tt.file)
		} else {
			t.Logf("result: got %d, wanted %d, for %s", result, tt.result, tt.file)
		}
	}
}
