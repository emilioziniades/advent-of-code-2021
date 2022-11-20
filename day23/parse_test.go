package day23_test

import (
	"reflect"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day23"
	"github.com/k0kubun/pp/v3"
)

func TestParseInitialState(t *testing.T) {
	tests := []struct {
		filename string
		expected day23.State
	}{
		{
			/*
				#############
				#...........#
				###B#C#B#D###
				  #A#D#C#A#
				  #########
			*/
			"example.txt",
			day23.State{
				{day23.Point{2, 3}, "B"},
				{day23.Point{2, 5}, "C"},
				{day23.Point{2, 7}, "B"},
				{day23.Point{2, 9}, "D"},
				{day23.Point{3, 3}, "A"},
				{day23.Point{3, 5}, "D"},
				{day23.Point{3, 7}, "C"},
				{day23.Point{3, 9}, "A"},
			},
		},
	}

	for _, tt := range tests {
		if got := day23.ParseState(tt.filename); !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("TestParseInitialState: file %v: got %v wanted %v", tt.filename, pp.Sprintln(got), pp.Sprintln(tt.expected))
		}
	}
}

func TestSortState(t *testing.T) {
	tests := []struct {
		input    day23.State
		expected day23.State
	}{
		{
			day23.State{
				{day23.Point{2, 3}, "B"},
				{day23.Point{1, 4}, "C"},
				{day23.Point{2, 7}, "B"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
			},
			day23.State{
				{day23.Point{1, 4}, "C"},
				{day23.Point{2, 3}, "B"},
				{day23.Point{2, 7}, "B"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
				{day23.Point{9, 9}, "X"},
			},
		},
	}

	for _, tt := range tests {
		got := tt.input
		day23.SortState(got[:])
		if got != tt.expected {
			t.Errorf("TestSortState: got != expected, got: %v, expected: %v", tt.input, tt.expected)
		}
	}

}
