package day21_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day21"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/21/input", "21.in")
	if err != nil {
		log.Fatal(err)
	}
}
func TestPlay(t *testing.T) {
	var tests = []struct {
		file          string
		deterministic int
		quantum       int
	}{
		{"21.ex", 739785, 444356092776315},
		{"21.in", 671580, 912857726749764},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		deterministic := day21.Play(in)
		quantum := day21.PlayDirac(in)

		format := "%s\n\tgot %d, wanted %d, for deterministic play\n\tgot %d wanted %d for quantum play\n"
		if deterministic != tt.deterministic || quantum != tt.quantum {
			t.Fatalf(format, tt.file, deterministic, tt.deterministic, quantum, tt.quantum)
		} else {
			t.Logf(format, tt.file, deterministic, tt.deterministic, quantum, tt.quantum)
		}
	}
}
