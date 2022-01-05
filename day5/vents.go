package day5

import (
	"fmt"
	"log"
	"strings"

	"github.com/emilioziniades/adventofcode2021/fetch"
)

type vent struct {
	x1, x2, y1, y2 int
}

type ventMap [][]int

func (v ventMap) String() string {
	s := ""
	for i := range v {
		s += fmt.Sprintf("%v\n", v[i])
	}
	return s
}

func (v ventMap) countTwos(size int) int {
	var n int
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if v[y][x] >= 2 {
				n++
			}
		}
	}
	return n
}

func MapVents(coords []string, diagonal bool) int {
	vents := parseCoords(coords)
	size := squareSize(vents)
	var vm ventMap = make([][]int, size)
	for i := range vm {
		vm[i] = make([]int, size)
	}

	mapVentsStraight(vents, vm)
	if diagonal {
		mapVentsDiagonal(vents, vm)
	}
	fmt.Println(vm)
	return vm.countTwos(size)
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
		var xS, xE, y int
		switch x1, y1, x2, y2 := v.x1, v.y1, v.x2, v.y2; {
		case x1 < x2 && y1 > y2:
			xS, xE = x1, x2
			y = y2
		case x1 < x2 && y1 < y2:
			xS, xE = x1, x2
			y = y1
		case x1 > x2 && y1 < y2:
			xS, xE = x2, x1
			y = y1
		case x1 > x2 && y1 > y2:
			xS, xE = x2, x1
			y = y2
		default:
			continue Loop
		}
		for x := xS; x <= xE; x++ {
			vm[y][x]++
			y++
		}
	}
}
func split(r rune) bool {
	return r == ',' || r == '-' || r == '>' || r == ' '
}
