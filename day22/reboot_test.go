package day22_test

import (
	"testing"

	"github.com/emilioziniades/adventofcode2021/day22"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/22/input", "22.in")
	if err != nil {
		panic(err)
	}
}
func TestReboot(t *testing.T) {
	var tests = []struct {
		file  string
		want  int
		limit float64
	}{
		{"22.si", 39, 50},
		// {"22.ex", 590784, 50},
		// {"22.ex2", 474140, 50},
		// {"22.in", 553201, 50},
		// {"22.ex2", 474140, math.MaxInt},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			panic(err)
		}

		got := day22.Reboot(in, tt.limit)
		format := "got %d, wanted %d for %s"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.file)
		} else {
			t.Logf(format, got, tt.want, tt.file)
		}
	}
}
