package day25_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day25"
	"github.com/emilioziniades/adventofcode2021/fetch"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/25/input", "input.txt")
	if err != nil {
		log.Fatal(err)
	}
}
func TestParseInput(t *testing.T) {
	tests := []struct {
		file      string
		wantFloor day25.SeaFloor
	}{
		{
			"simple.txt",
			day25.SeaFloor{
				day25.Point{0, 3}: day25.East,
				day25.Point{0, 4}: day25.East,
				day25.Point{0, 5}: day25.East,
				day25.Point{0, 6}: day25.East,
				day25.Point{0, 7}: day25.East,
			},
		},
		{
			"example.txt",
			day25.SeaFloor{
				day25.Point{0, 0}: day25.South,
				day25.Point{8, 9}: day25.East,
				day25.Point{4, 4}: day25.South,
				day25.Point{4, 5}: day25.South,
			},
		},
	}
	for _, tt := range tests {
		state := day25.ParseInput(tt.file)

		for wantPoint, wantType := range tt.wantFloor {
			gotType, ok := state.F[wantPoint]
			if !ok {
				t.Errorf("point not found %v", wantPoint)
			}

			if gotType != wantType {
				t.Errorf("cucumber at point %v not correct type, wanted %v, got %v", wantPoint, wantType, gotType)
			}
		}
	}
}

func TestStep(t *testing.T) {
	tests := []struct {
		initial string
		nSteps  int
		want    string
	}{
		{
			"simple.txt",
			1,
			"...>>>>.>..\n",
		},
		{
			"simple.txt",
			2,
			"...>>>.>.>.\n",
		},
		{
			"simple_2.txt",
			1,
			string(
				`..........
.>........
..v....v>.
..........
`),
		},
	}

	for _, tt := range tests {
		state := day25.ParseInput(tt.initial)
		for i := 1; i <= tt.nSteps; i++ {
			state.StepBoth()
		}

		if got := state.String(); tt.want != got {
			t.Errorf("\n%v:\nwant\n%v\ngot\n%v\n", tt.initial, tt.want, got)
		}
	}
}

func TestStepUntil(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{
			"example.txt",
			58,
		},
		{
			"input.txt",
			504,
		},
	}

	for _, tt := range tests {
		state := day25.ParseInput(tt.filename)
		got := day25.StepUntilEnd(state)
		if got != tt.want {
			t.Errorf("got %v, wanted %v, for %v", got, tt.want, tt.filename)
		}
	}
}
