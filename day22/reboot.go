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

func newCuboid(s, e point, on bool) cuboid {
	if side := e.x - s.x; e.y-s.x != side || e.z-s.z != side {
		panic("not a cuboid")
	}
	return cuboid{s, e, on}
}

func c1StartInC2(c1, c2 cuboid) bool {
	return (c1.start.x >= c2.start.x && c1.start.x <= c2.end.x) && (c1.start.y >= c2.start.y && c1.start.y <= c2.end.y) && (c1.start.z >= c2.start.z && c1.start.z <= c2.end.z)
}

func c2StartInC1(c1, c2 cuboid) bool {
	return (c1.end.x >= c2.start.x && c1.end.x <= c2.end.x) && (c1.end.y >= c2.start.y && c1.end.y <= c2.end.y) && (c1.end.z >= c2.start.z && c1.end.z <= c2.end.z)
}

func isOverlap(c1, c2 cuboid) bool {
	//does one cube have a start or end point that lies within the other cube?
	return c1StartInC2(c1, c2) || c2StartInC1(c1, c2)
}

func (c cuboid) Volume() int {
	return (c.end.x - c.start.x + 1) * (c.end.y - c.start.y + 1) * (c.end.z - c.start.z + 1)
}

func Reboot(input []string, bounds float64) int {
	reactor := make(map[point]bool)
	cuboids := stack.New[cuboid]()
	// cuboids := make([]cuboid, 0)
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

			currCuboid := newCuboid(point{xS, yS, zS}, point{xE, yE, zE}, on)
			cuboids.Push(currCuboid)
			fmt.Println(currCuboid)
			fmt.Println(line)
			fmt.Println(currCuboid.Volume())
			continue

			for x := xS; x <= xE; x++ {
				for y := yS; y <= yE; y++ {
					for z := zS; z <= zE; z++ {
						fmt.Printf("%d,%d,%d\n", x, y, z)
						reactor[point{x, y, z}] = on

					}
				}
			}
		} else {
			panic("parsing error")
		}
	}

	for _, c1 := range cuboids {
		for _, c2 := range cuboids {
			if isOverlap(c1, c2) {
				// fmt.Println(c1, c2)
			}
		}
	}

	//try and split first two cuboids
	fmt.Println(cuboids)
	c1, c2 := cuboids.PopLeft(), cuboids.PopLeft()
	fmt.Println(c1, c2)
	fmt.Println(isOverlap(c1, c2))
	if c1StartInC2(c1, c2) {
		fmt.Println("c1 start in c2")
	} else if c2StartInC1(c1, c2) {
		fmt.Println("c2 start in c1")
		nc1 := cuboid{c1.start, c2.start, c1.on}
		nc2 := cuboid{c2.start, c1.end, false}
		nc3 := cuboid{c1.end, c2.end, false}
		fmt.Println(nc1, nc2, nc3)
	} else {
	}
	return countOn(reactor)
}

func countOn(r map[point]bool) (on int) {
	for _, o := range r {
		if o == true {
			on++
		}
	}
	return
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

// when two cuboids overlap, they can be split into four distinct cuboids and treated separately
// func Split(c1, c2 cuboid) (nc1, nc2, nc3, nc4 cuboid) {
// }
