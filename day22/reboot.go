package day22

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/emilioziniades/adventofcode2021/stack"
)

type point struct {
	x, y, z int
}

type cuboid struct {
	start, end point
	on         bool
}

func (c cuboid) Volume() int {
	return (c.end.x - c.start.x + 1) * (c.end.y - c.start.y + 1) * (c.end.z - c.start.z + 1)
}

func Reboot(input []string, bounds float64) int {
	// reactor := make(map[point]bool)
	reactor := make(map[cuboid]bool)
	cuboids := inputToCuboids(input, bounds)

	/*
		The logic for adding cubes to the reactor map is as follows:
		For every cuboid from the input:
			If not intersected with any other existing cuboid in reactor:
				Add to reactor
			Else if there is intersection:
				Split two intersecting cuboids into up to 7 child cuboids, with appropriate on/off status
				Then check if those children cuboids intersect with existing cuboids in reactor:
					If no intersect:
						Add to reactor
					If intersect:
						Recursively do the above procedure

	*/

	//try and split first two cuboids
	fmt.Println(cuboids)
	c1, c2 := cuboids.PopLeft(), cuboids.PopLeft()
	fmt.Println(isOverlap(c1, c2))
	Split(c1, c2)
	return countOnCuboid(reactor)
}

func isOverlap(a, b cuboid) bool {
	//does one cube have a start or end point that lies within the other cube?
	return AStartInB(a, b) || AStartInB(b, a)
}

func AStartInB(a, b cuboid) bool {
	return (a.start.x >= b.start.x && a.start.x <= b.end.x) && (a.start.y >= b.start.y && a.start.y <= b.end.y) && (a.start.z >= b.start.z && a.start.z <= b.end.z)
}

// when two cuboids overlap, they can be split into four distinct cuboids and treated separately
func Split(c1, c2 cuboid) []cuboid {
	if AStartInB(c1, c2) {
		return split(c2, c1, true)
	} else if AStartInB(c2, c1) {
		return split(c1, c2, false)
	} else {
		return make([]cuboid, 0)
	}
}

// for now, assuming that one cuboid does not contain the other cuboid
func split(a, b cuboid, inverted bool) []cuboid {
	res := make([]cuboid, 0)
	fmt.Println(a, a.Volume())
	fmt.Println(b, b.Volume())
	// two chunks on either side with no overlap
	nc1 := cuboid{point{a.start.x, a.start.y, a.start.z}, point{b.start.x - 1, a.end.y, a.end.z}, a.on}
	nc2 := cuboid{point{a.end.x, b.start.y, b.start.z}, point{b.end.x - 1, b.end.y, b.end.z}, b.on}

	// overlapping section
	overlapOn := false
	if inverted {
		overlapOn = a.on
	} else {
		overlapOn = b.on
	}
	nc3 := cuboid{b.start, a.end, overlapOn}

	// four remaining chunks left, below, above and right of the overlapping section
	nc4 := cuboid{point{b.start.x, a.start.y, a.start.z}, point{a.end.x - 1, b.start.y, a.end.z}, a.on}
	nc5 := cuboid{point{b.start.x, b.start.y, a.start.z}, point{a.end.x - 1, a.end.y, b.start.z}, a.on}
	nc6 := cuboid{point{b.start.x, b.start.y, a.end.z}, point{a.end.x - 1, a.end.y, b.end.z}, b.on}
	nc7 := cuboid{point{b.start.x, a.end.y, b.start.z}, point{a.end.x - 1, b.end.y, b.end.z}, b.on}

	res = append(res, nc1, nc2, nc3, nc4, nc5, nc6, nc7)
	// RenderCuboids(res)

	sum := 0
	for _, c := range res {
		fmt.Println(c, c.Volume())
		sum += c.Volume()
	}
	fmt.Println(sum)
	fmt.Println(res)
	return res
}

// countOnCuboid counts the number of unit cubes that are on in the reactor map. It assumes that there are no overlapping cuboids, and such overlaps would be handled before insertion into reactor map
func countOnCuboid(r map[cuboid]bool) (on int) {
	for c, o := range r {
		if o {
			on += c.Volume()
		}
	}
	return
}

func inputToCuboids(input []string, bounds float64) stack.Stack[cuboid] {
	cuboids := stack.New[cuboid]()
	re := regexp.MustCompile(`([a-z]+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+)`)
	for _, line := range input {
		if re.MatchString(line) {
			match := re.FindStringSubmatch(line)
			on := false
			if match[1] == "on" {
				on = true
			}
			xS, xE := intInRange(match[2], bounds, true), intInRange(match[3], bounds, false)
			yS, yE := intInRange(match[4], bounds, true), intInRange(match[5], bounds, false)
			zS, zE := intInRange(match[6], bounds, true), intInRange(match[7], bounds, false)

			currCuboid := cuboid{point{xS, yS, zS}, point{xE, yE, zE}, on}
			cuboids.Push(currCuboid)

		} else {
			panic("parsing error")
		}
	}
	return cuboids
}

func intInRange(s string, r float64, start bool) int {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	if start {
		return int(math.Max(f, -1*r))
	} else {
		return int(math.Min(f, r))
	}
}
