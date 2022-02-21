package day22

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/emilioziniades/adventofcode2021/stack"
)

func isOverlap(a, b Cuboid) bool {
	//does one cube have a start or end Point that lies within the other cube?
	return isAStartInB(a, b) || isAStartInB(b, a)
}

func isAStartInB(a, b Cuboid) bool {
	return (a.start.x >= b.start.x && a.start.x <= b.end.x) && (a.start.y >= b.start.y && a.start.y <= b.end.y) && (a.start.z >= b.start.z && a.start.z <= b.end.z)

}

func isAEndInB(a, b Cuboid) bool {
	return (a.end.x >= b.start.x && a.end.x <= b.end.x) && (a.end.y >= b.start.y && a.end.y <= b.end.y) && (a.end.z >= b.start.z && a.end.z <= b.end.z)
}

// countOn counts the number of unit cubes that are on in the reactor map. It assumes that there are no overlapping cuboids, and such overlaps would be handled before insertion into reactor map
func countOn(r map[Cuboid]bool) (on int) {
	for c, o := range r {
		// fmt.Println(c, c.Volume(), o)
		if o {
			on += c.Volume()
		}
	}
	return
}

func printPoints(c Cuboid) {
	for x := c.start.x; x <= c.end.x; x++ {
		for y := c.start.y; y <= c.end.y; y++ {
			for z := c.start.z; z <= c.end.z; z++ {
				fmt.Printf("%d,%d,%d\n", x, y, z)
			}
		}
	}
}

func printAllPoints(r map[Cuboid]bool) {
	for k := range r {
		printPoints(k)
	}
}

func inputToCuboids(input []string, bounds int) stack.Stack[Cuboid] {
	cuboids := stack.New[Cuboid]()
	boundary := NewCuboid(-1*bounds, -1*bounds, -1*bounds, bounds, bounds, bounds, true)
	re := regexp.MustCompile(`([a-z]+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+)`)
	for _, line := range input {
		if re.MatchString(line) {
			match := re.FindStringSubmatch(line)
			on := false
			if match[1] == "on" {
				on = true
			}

			fmt.Println(line)

			xS, xE := toInt(match[2]), toInt(match[3])
			yS, yE := toInt(match[4]), toInt(match[5])
			zS, zE := toInt(match[6]), toInt(match[7])

			currCuboid := NewCuboid(xS, yS, zS, xE, yE, zE, on)

			if isOverlap(*boundary, *currCuboid) {

				cuboids.Push(boundCuboid(*currCuboid, bounds))
			}

		} else {
			panic("parsing error")
		}
	}
	return cuboids
}

func intInRange(n int, b int, start bool) int {
	num := float64(n)
	lower, upper := float64(-1*b), float64(b)
	if start {
		return int(math.Max(num, lower))
	} else {
		return int(math.Min(num, upper))
	}
}

func toInt(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}

func boundCuboid(c Cuboid, bounds int) Cuboid {
	xS, xE := intInRange(c.start.x, bounds, true), intInRange(c.end.x, bounds, false)
	yS, yE := intInRange(c.start.y, bounds, true), intInRange(c.end.y, bounds, false)
	zS, zE := intInRange(c.start.z, bounds, true), intInRange(c.end.z, bounds, false)
	return *NewCuboid(xS, yS, zS, xE, yE, zE, c.on)
}
