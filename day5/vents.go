package day5

import (
	"log"
	"strings"

	"github.com/emilioziniades/adventofcode2021/fetch"
)

type vent struct {
	x1, y1, x2, y2 int
}

type ventMap [][]int

func (vm ventMap) String() string {
	s := ""
	for i := range vm {
		ss := fetch.IntToStringSlice(vm[i])
		s += strings.Replace(strings.Join(ss, ""), "0", ".", -1)
		s += "\n"
	}
	return s
}

func (vm ventMap) countTwos() int {
	var n int
	for y := 0; y < len(vm); y++ {
		for x := 0; x < len(vm); x++ {
			if vm[y][x] >= 2 {
				n++
			}
		}
	}
	return n
}

func makeVentMap(size int) ventMap {
	var vm ventMap = make([][]int, size)
	for i := range vm {
		vm[i] = make([]int, size)
	}
	return vm
}

func MapVents(coords []string, diagonal bool) int {
	vents := parseCoords(coords)
	size := squareSize(vents)
	vm := makeVentMap(size)

	mapVentsStraight(vents, vm)
	if diagonal {
		mapVentsDiagonal(vents, vm)
	}
	return vm.countTwos()
}

func mapVentsStraight(vents []vent, vm ventMap) {
Loop:
	for _, v := range vents {
		var start, end, constant int
		var yconst bool
		switch x1, y1, x2, y2 := v.x1, v.y1, v.x2, v.y2; {
		case x1 == x2 && y2 > y1:
			start, end, constant = y1, y2, x1
			yconst = false
		case x1 == x2 && y1 > y2:
			start, end, constant = y2, y1, x1
			yconst = false
		case y1 == y2 && x2 > x1:
			start, end, constant = x1, x2, y1
			yconst = true
		case y1 == y2 && x1 > x2:
			start, end, constant = x2, x1, y1
			yconst = true
		default:
			continue Loop
		}

		for i := start; i <= end; i++ {
			if yconst {
				vm[constant][i]++
			} else {
				vm[i][constant]++
			}
		}
	}
}

func mapVentsDiagonal(vents []vent, vm ventMap) {
Loop:
	for _, v := range vents {
		var xS, xE, yS, yE int
		switch x1, y1, x2, y2 := v.x1, v.y1, v.x2, v.y2; {
		case x1 < x2 && y1 > y2:
			xS, xE = x1, x2
			yS, yE = y1, y2
		case x1 < x2 && y1 < y2:
			xS, xE = x1, x2
			yS, yE = y1, y2
		case x1 > x2 && y1 < y2:
			xS, xE = x2, x1
			yS, yE = y2, y1
		case x1 > x2 && y1 > y2:
			xS, xE = x2, x1
			yS, yE = y2, y1
		default:
			continue Loop
		}

		y := yS
		for x := xS; x <= xE; x++ {
			vm[y][x]++
			if yS < yE {
				y++
			} else {
				y--
			}
		}
	}
}

func parseCoords(coords []string) []vent {
	vents := make([]vent, 0)
	for _, coord := range coords {
		cString := strings.FieldsFunc(coord, split)
		c, err := fetch.StringToIntSlice(cString)
		if err != nil {
			log.Fatalf("MapVents: parseCoords: %s", err)
		}
		v := vent{x1: c[0], y1: c[1], x2: c[2], y2: c[3]}
		vents = append(vents, v)
	}
	return vents
}

func squareSize(vents []vent) int {
	var x, y int
	for _, v := range vents {
		switch {
		case v.x1 > x:
			x = v.x1
		case v.x2 > x:
			x = v.x2
		case v.y1 > y:
			y = v.y1
		case v.y2 > y:
			y = v.y2
		}
	}

	if x >= y {
		return x + 1
	} else {
		return y + 1
	}
}

func split(r rune) bool {
	return r == ',' || r == '-' || r == '>' || r == ' '
}
