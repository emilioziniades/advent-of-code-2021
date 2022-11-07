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

func TestHomeButMustMakeSpace(t *testing.T) {
	state := day23.ParseInitialState("example.txt")
	tests := []struct {
		point    day23.Point
		podType  rune
		expected bool
	}{
		// bottom row - never needs to make space
		{
			day23.Point{3, 3},
			'A',
			false,
		},
		// bottom row
		{
			day23.Point{3, 5},
			'D',
			false,
		},
		// top row, but not home
		{
			day23.Point{2, 5},
			'C',
			false,
		},
		// top row and is home
		{
			day23.Point{2, 9},
			'D',
			true,
		},
	}

	for _, tt := range tests {
		if got := day23.HomeButMustMakeSpace(tt.point, tt.podType, &state); got != tt.expected {
			t.Errorf("TestHomeButMustMakeSpace: got %v wanted %v for %#v %#v", got, tt.expected, tt.point, string(tt.podType))
		}
	}
}

func TestPodNextPositionsAndCosts(t *testing.T) {
	state := day23.ParseInitialState("example.txt")
	tests := []struct {
		podPosition       day23.Point
		podType           rune
		expectedPositions map[day23.Point]int
	}{
		// nowhere to go
		{
			day23.Point{3, 9},
			'A',
			map[day23.Point]int{},
		},
		// also nowhere to go
		{
			day23.Point{3, 7},
			'C',
			map[day23.Point]int{},
		},
		// already home, can't go anywhere
		{
			day23.Point{3, 3},
			'A',
			map[day23.Point]int{},
		},
		// "home" but has to make space
		{
			day23.Point{2, 9},
			'D',
			map[day23.Point]int{
				{1, 1}:  9000,
				{1, 2}:  8000,
				{1, 4}:  6000,
				{1, 6}:  4000,
				{1, 8}:  2000,
				{1, 10}: 2000,
				{1, 11}: 3000,
			},
		},
		// can go to all hallway positions
		{
			day23.Point{2, 3},
			'B',
			map[day23.Point]int{
				{1, 1}:  30,
				{1, 2}:  20,
				{1, 4}:  20,
				{1, 6}:  40,
				{1, 8}:  60,
				{1, 10}: 80,
				{1, 11}: 90,
			},
		},
	}

	for _, tt := range tests {
		got := day23.GetPodNextPositionsAndCosts(tt.podPosition, tt.podType, &state)
		for gotPoint, gotCost := range got {
			// is this point in expected
			expectedCost, ok := (tt.expectedPositions)[gotPoint]
			if ok && expectedCost != gotCost {
				t.Errorf("TestPodNextPositionAndCosts: costs not equal, got %v, wanted %v for point %#v", gotCost, expectedCost, gotPoint)
			}
			if !ok {
				t.Errorf("TestPodNextPositionAndCosts: position not found, point %#v", gotPoint)
			}
			delete(tt.expectedPositions, gotPoint)
		}

		if len(tt.expectedPositions) != 0 {
			t.Errorf("TestPodNextPositionsAndCosts: not all positions found: %v", tt.expectedPositions)
		}
	}

}

func TestDjikstra(t *testing.T) {
	day23.Djikstra("example.txt")
}
