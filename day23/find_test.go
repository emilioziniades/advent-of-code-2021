package day23_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day23"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/k0kubun/pp/v3"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/23/input", "input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

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
		podType      string
		expectedDist int
		expectedCost int
	}{
		{
			// no movement
			day23.Point{1, 1},
			day23.Point{1, 1},
			"A",
			0,
			0,
		},
		{
			// one across
			day23.Point{1, 1},
			day23.Point{1, 2},
			"B",
			1,
			10,
		},
		{
			// all the way across the hall
			day23.Point{1, 1},
			day23.Point{1, 11},
			"C",
			10,
			1000,
		},
		{
			// down into room
			day23.Point{1, 1},
			day23.Point{3, 3},
			"D",
			4,
			4000,
		},
		{
			// one room to another
			day23.Point{3, 3},
			day23.Point{3, 9},
			"A",
			10,
			10,
		},
		{
			// upper room to hallway spot
			day23.Point{2, 3},
			day23.Point{1, 7},
			"B",
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
		p      day23.Pod
		inHall bool
	}{
		{
			day23.Pod{day23.Point{1, 1}, "X"},
			true,
		},
		{
			day23.Pod{day23.Point{1, 8}, "X"},
			true,
		},
		{
			day23.Pod{day23.Point{2, 7}, "X"},
			false,
		},
		{
			day23.Pod{day23.Point{3, 9}, "X"},
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
		p      day23.Pod
		inHome bool
	}{
		{
			// in hall
			day23.Pod{
				day23.Point{1, 1},
				"A",
			},
			false,
		},
		{
			// B in A home
			day23.Pod{
				day23.Point{3, 3},
				"B",
			},
			false,
		},
		{
			// C in C home
			day23.Pod{
				day23.Point{2, 7},
				"C",
			},
			true,
		},
		{
			// D in C home
			day23.Pod{
				day23.Point{3, 7},
				"D",
			},
			false,
		},
	}

	for _, tt := range tests {
		if got := tt.p.IsHome(); got != tt.inHome {
			t.Errorf("TestInHallway: got %v wanted %v", got, tt.inHome)
		}
	}
}

func TestHomeButMustMakeSpace(t *testing.T) {
	state := day23.ParseState("example.txt", false)
	tests := []struct {
		pod      day23.Pod
		expected bool
	}{
		// bottom row - never needs to make space
		{
			day23.Pod{
				day23.Point{3, 3},
				"A",
			},
			false,
		},
		// bottom row
		{
			day23.Pod{
				day23.Point{3, 5},
				"D",
			},
			false,
		},
		// top row, but not home
		{
			day23.Pod{
				day23.Point{2, 5},
				"C",
			},
			false,
		},
		// top row and is home
		{
			day23.Pod{
				day23.Point{2, 9},
				"D",
			},
			true,
		},
	}

	for _, tt := range tests {
		if got := tt.pod.HomeButMustMakeSpace(state); got != tt.expected {
			t.Errorf("TestHomeButMustMakeSpace: got %v wanted %v for %#v", got, tt.expected, tt.pod)
		}
	}
}

func TestPodNextPositionsAndCosts(t *testing.T) {
	type testCase struct {
		pod      day23.Pod
		expected map[day23.Pod]int
	}
	tests := []struct {
		filename string
		cases    []testCase
	}{
		{
			"example.txt",
			[]testCase{
				// nowhere to go
				{
					day23.Pod{
						day23.Point{3, 9},
						"A",
					},
					map[day23.Pod]int{},
				},
				// also nowhere to go
				{
					day23.Pod{
						day23.Point{3, 7},
						"C",
					},
					map[day23.Pod]int{},
				},
				// already home, can't go anywhere
				{
					day23.Pod{
						day23.Point{3, 3},
						"A",
					},
					map[day23.Pod]int{},
				},
				// "home" but has to make space
				{
					day23.Pod{
						day23.Point{2, 9},
						"D",
					},
					map[day23.Pod]int{
						{day23.Point{1, 1}, "D"}:  9000,
						{day23.Point{1, 2}, "D"}:  8000,
						{day23.Point{1, 4}, "D"}:  6000,
						{day23.Point{1, 6}, "D"}:  4000,
						{day23.Point{1, 8}, "D"}:  2000,
						{day23.Point{1, 10}, "D"}: 2000,
						{day23.Point{1, 11}, "D"}: 3000,
					},
				},
				// can go to all hallway positions
				{
					day23.Pod{
						day23.Point{2, 3},
						"B",
					},
					map[day23.Pod]int{
						{day23.Point{1, 1}, "B"}:  30,
						{day23.Point{1, 2}, "B"}:  20,
						{day23.Point{1, 4}, "B"}:  20,
						{day23.Point{1, 6}, "B"}:  40,
						{day23.Point{1, 8}, "B"}:  60,
						{day23.Point{1, 10}, "B"}: 80,
						{day23.Point{1, 11}, "B"}: 90,
					},
				},
			},
		},
		{
			"example2.txt",
			[]testCase{
				{
					// from starting position, it can go home or into hall
					day23.Pod{
						day23.Point{2, 5},
						"C",
					},
					map[day23.Pod]int{
						// home
						{day23.Point{2, 7}, "C"}: 400,
						// all possible hallway positions
						{day23.Point{1, 6}, "C"}:  200,
						{day23.Point{1, 8}, "C"}:  400,
						{day23.Point{1, 10}, "C"}: 600,
						{day23.Point{1, 11}, "C"}: 700,
					},
				},
			},
		},
		{
			"example3.txt",
			[]testCase{
				{
					// from hall into home
					day23.Pod{
						day23.Point{1, 4},
						"B",
					},
					map[day23.Pod]int{
						{day23.Point{3, 5}, "B"}: 30, // home
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		state := day23.ParseState(testCase.filename, false)
		for _, test := range testCase.cases {
			got := day23.GetPodNextPositionsAndCosts(test.pod, state)
			if !reflect.DeepEqual(got, test.expected) {

				t.Errorf("All next positions not found for pod %#v", test.pod)
				pp.Println(got)
				fmt.Println()
				pp.Println(test.expected)
			}
		}
	}

}

func TestDjikstra(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{
			"example.txt",
			12521,
		},
		{
			"input.txt",
			19059,
		},
	}

	for _, tt := range tests {
		got := day23.Djikstra(tt.filename, false)
		if got != tt.want {
			t.Errorf("TestDjikstra: wanted %v, got %v", tt.want, got)
		}
	}
}

func TestBigDjikstra(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{
			"example.txt",
			44169,
		},
		{
			"input.txt",
			48541,
		},
	}

	for _, tt := range tests {
		got := day23.Djikstra(tt.filename, true)
		if got != tt.want {
			t.Errorf("TestBigDjikstra: wanted %v, got %v", tt.want, got)
		}
	}
}
