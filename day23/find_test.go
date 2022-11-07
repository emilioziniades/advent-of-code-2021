package day23_test

import (
	"reflect"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day23"
)

func TestAbs(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{-10, 10},
		{0, 0},
		{42, 42},
		{-420, 420},
	}

	for _, tt := range tests {
		if got := day23.Abs(tt.input); got != tt.expected {
			t.Errorf("TestAbs: wanted %v got %v", tt.expected, got)
		}
	}

}

/*
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########

(1,1) (1,2) (1,3) (1,4) (1,5) (1,6) (1,7) (1,8) (1,9) (1,10) (1,11)
            (2,3)       (2,5)       (2,7)       (2,9)
            (3,3)       (3,5)       (3,7)       (3,9)

*/

func TestDistanceAndCost(t *testing.T) {
	var tests = []struct {
		a            day23.Point
		b            day23.Point
		podType      rune
		expectedDist int
		expectedCost int
	}{
		{
			// no movement
			day23.Point{1, 1},
			day23.Point{1, 1},
			'A',
			0,
			0,
		},
		{
			// one across
			day23.Point{1, 1},
			day23.Point{1, 2},
			'B',
			1,
			10,
		},
		{
			// all the way across the hall
			day23.Point{1, 1},
			day23.Point{1, 11},
			'C',
			10,
			1000,
		},
		{
			// down into room
			day23.Point{1, 1},
			day23.Point{3, 3},
			'D',
			4,
			4000,
		},
		{
			// one room to another
			day23.Point{3, 3},
			day23.Point{3, 9},
			'A',
			10,
			10,
		},
		{
			// upper room to hallway spot
			day23.Point{2, 3},
			day23.Point{1, 7},
			'B',
			5,
			50,
		},
	}

	for _, tt := range tests {
		if got := day23.Distance(tt.a, tt.b); got != tt.expectedDist {
			t.Errorf("TestDistance: wanted %v, got %v", tt.expectedDist, got)
		}
		if got := day23.DistanceCost(tt.a, tt.b, tt.podType); got != tt.expectedCost {
			t.Errorf("TestCost: wanted %v, got %v", tt.expectedCost, got)
		}
	}
}

func TestInHallway(t *testing.T) {
	var tests = []struct {
		p      day23.Point
		inHall bool
	}{
		{
			day23.Point{1, 1},
			true,
		},
		{
			day23.Point{1, 8},
			true,
		},
		{
			day23.Point{2, 7},
			false,
		},
		{
			day23.Point{3, 9},
			false,
		},
	}

	for _, tt := range tests {
		if got := tt.p.InHallway(); got != tt.inHall {
			t.Errorf("TestInHallway: got %v wanted %v", got, tt.inHall)
		}
	}
}

func TestIsHome(t *testing.T) {
	var tests = []struct {
		p       day23.Point
		podType rune
		inHome  bool
	}{
		{
			// in hall
			day23.Point{1, 1},
			'A',
			false,
		},
		{
			// B in A home
			day23.Point{3, 3},
			'B',
			false,
		},
		{
			// C in C home
			day23.Point{2, 7},
			'C',
			true,
		},
		{
			// D in C home
			day23.Point{3, 7},
			'D',
			false,
		},
	}

	for _, tt := range tests {
		if got := tt.p.IsHome(tt.podType); got != tt.inHome {
			t.Errorf("TestInHallway: got %v wanted %v", got, tt.inHome)
		}
	}
}

func TestCopy(t *testing.T) {
	state := &day23.State{day23.Point{1, 1}: 'A'}
	stateCopy := state.Copy()

	if !reflect.DeepEqual(*state, *stateCopy) {
		t.Errorf("TestCopy: copy and original not same")
	}

	if state == stateCopy {
		t.Errorf("TestCopy: copy and original share same pointer")
	}
}

func TestCopyWithout(t *testing.T) {
	state := &day23.State{
		day23.Point{1, 1}: 'A',
		day23.Point{2, 2}: 'B',
		day23.Point{3, 3}: 'C',
	}
	stateCopy := state.CopyWithout(day23.Point{1, 1})

	if _, ok := (*stateCopy)[day23.Point{1, 1}]; ok {
		t.Errorf("TestCopyWithout: {1,1}: A still in state")

	}

	if podType, ok := (*stateCopy)[day23.Point{2, 2}]; podType != 'B' || !ok {
		t.Errorf("TestCopyWithout: Missing {2,2}: B")

	}

	if podType, ok := (*stateCopy)[day23.Point{3, 3}]; podType != 'C' || !ok {
		t.Errorf("TestCopyWithout: Missing {3,3}: C")

	}
}
