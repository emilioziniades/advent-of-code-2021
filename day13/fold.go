package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type instruction struct {
	vertical bool
	line     int
}

func Fold(input []string, once bool) int {
	pts, insts := inputToPointsAndInstructions(input)
	paper := make(map[point]bool)

	for _, pt := range pts {
		paper[pt] = true
	}

	for i := range insts {
		for pt := range paper {
			curr := pt
			delete(paper, pt)
			folded := fold(curr, insts[i])
			paper[folded] = true
		}

		if once {
			break
		}
	}

	if !once {
		printPaper(paper)
	}
	return len(paper)
}

func fold(p point, i instruction) point {
	shift, fX, fY := float64(i.line), float64(p.x), float64(p.y)
	switch i.vertical {
	case true:
		newX := shift - math.Abs(fX-shift)
		return point{int(newX), p.y}
	case false:
		newY := shift - math.Abs(fY-shift)
		return point{p.x, int(newY)}
	default:
		return p
	}
}

func inputToPointsAndInstructions(input []string) ([]point, []instruction) {
	points := make([]point, 0)
	instructions := make([]instruction, 0)
	var pointsDone bool

	for _, e := range input {

		if e == "" {
			pointsDone = true
			continue
		}

		if pointsDone {
			inst := strings.Split(strings.TrimPrefix(e, "fold along "), "=")
			line, _ := strconv.Atoi(inst[1])
			instructions = append(instructions, instruction{inst[0] == "x", line})
			continue
		}
		nums := strings.Split(e, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		points = append(points, point{x, y})
	}
	return points, instructions

}

func findWidthHeight(paper map[point]bool) (int, int) {
	w, h := 0, 0
	for pt := range paper {
		if pt.x > w {
			w = pt.x
		}
		if pt.y > h {
			h = pt.y
		}
	}
	return w, h
}
func printPaper(paper map[point]bool) {
	w, h := findWidthHeight(paper)
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			if paper[point{x, y}] {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
