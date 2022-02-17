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

func TestSplit(t *testing.T) {
	var tests = []struct {
		c1, c2 day22.Cuboid
		want   int
	}{
		{day22.Cuboid{day22.Point{10, 10, 10}, day22.Point{12, 12, 12}, true}, day22.Cuboid{day22.Point{11, 11, 11}, day22.Point{13, 13, 13}, true}, 46},
		{day22.Cuboid{day22.Point{11, 11, 11}, day22.Point{13, 13, 13}, true}, day22.Cuboid{day22.Point{10, 10, 10}, day22.Point{12, 12, 12}, true}, 46},
		{day22.Cuboid{day22.Point{10, 10, 10}, day22.Point{13, 13, 13}, true}, day22.Cuboid{day22.Point{12, 12, 12}, day22.Point{15, 15, 15}, true}, 120},
		{day22.Cuboid{day22.Point{9, 9, 9}, day22.Point{11, 11, 11}, true}, day22.Cuboid{day22.Point{10, 10, 10}, day22.Point{10, 12, 12}, true}, 32},
		{day22.Cuboid{day22.Point{11, 9, 9}, day22.Point{11, 11, 11}, true}, day22.Cuboid{day22.Point{11, 10, 10}, day22.Point{12, 10, 12}, true}, 13},
	}
	for _, tt := range tests {
		children := day22.Split(tt.c1, tt.c2)
		count := 0
		for _, c := range children {
			if c.Volume() < 0 {
				continue
			}
			count += c.Volume()
		}
		format := "wanted %d, got %d, for day22.Cuboids %v and %v"
		if count != tt.want {
			t.Fatalf(format, tt.want, count, tt.c1, tt.c2)
		} else {
			t.Logf(format, tt.want, count, tt.c1, tt.c2)
		}
	}
}
