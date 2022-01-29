package day19_test

import (
	"testing"

	"github.com/emilioziniades/adventofcode2021/day19"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func TestCountBeacons(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{"19.ex", 79},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			t.Fatal(err)
		}
		got := day19.CountBeacons(in)
		format := "got %d, wanted %d, for %s"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.file)
		} else {
			t.Logf(format, got, tt.want, tt.file)
		}
	}
}
