package day25

import (
	"log"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

type (
	Point struct {
		X, Y int
	}

	Cucumber int

	SeaFloor map[Point]Cucumber

	FloorState struct {
		maxX, maxY int
		F          SeaFloor
	}
)

const (
	South Cucumber = iota
	East
)

func StepUntilEnd(fs FloorState) int {
	for i := 1; ; i++ {
		if n := fs.StepBoth(); n == 0 {
			return i
		}
	}

}

func (fs FloorState) StepBoth() int {
	return fs.Step(East) + fs.Step(South)

}

func (fs FloorState) Step(cucType Cucumber) int {
	nMoves := 0
	newFloor := make(SeaFloor)

	// move cucs in newFloor
	for pt, cuc := range fs.F {
		if cuc == cucType {
			nextPt := fs.NextPosition(pt, cuc)
			if fs.F.IsEmpty(nextPt) {
				newFloor[nextPt] = cuc
				nMoves++
			} else {
				newFloor[pt] = cuc
			}
		}
	}

	fs.Clear(cucType)

	for pt, cuc := range newFloor {
		if cuc == cucType {
			fs.F[pt] = cuc
		}
	}

	return nMoves

}

func (fs *FloorState) Clear(cuc Cucumber) {
	for pt, c := range fs.F {
		if cuc == c {
			delete(fs.F, pt)
		}
	}
}

func (fs FloorState) NextPosition(pt Point, cuc Cucumber) Point {
	switch cuc {
	case East:
		if pt.Y == fs.maxY {
			return Point{pt.X, 0}
		} else {
			return Point{pt.X, pt.Y + 1}
		}
	case South:
		if pt.X == fs.maxX {
			return Point{0, pt.Y}
		} else {
			return Point{pt.X + 1, pt.Y}
		}
	default:
		panic("Unrecognized cucumber type")
	}

}

func (f SeaFloor) IsEmpty(pt Point) bool {
	_, ok := f[pt]
	return !ok

}

func ParseInput(filename string) FloorState {
	floor := make(SeaFloor)
	stringSlice, err := parse.FileToStringSlice(filename)
	if err != nil {
		log.Fatal(err)
	}

	for row, str := range stringSlice {
		for col, char := range str {
			pt := Point{row, col}
			switch string(char) {
			case ">":
				floor[pt] = East
			case "v":
				floor[pt] = South
			}
		}
	}

	state := FloorState{
		maxX: len(stringSlice) - 1,
		maxY: len(stringSlice[0]) - 1,
		F:    floor}

	return state

}

func (c Cucumber) String() string {
	switch c {
	case South:
		return "v"
	case East:
		return ">"
	default:
		return "."
	}
}

func (fs FloorState) String() string {
	s := strings.Builder{}

	for i := 0; i <= fs.maxX; i++ {
		for j := 0; j <= fs.maxY; j++ {
			pt := Point{i, j}
			cuc, ok := fs.F[pt]
			if !ok {
				s.WriteString(".")
			} else {
				s.WriteString(cuc.String())
			}
		}
		s.WriteString("\n")
	}
	return s.String()

}
