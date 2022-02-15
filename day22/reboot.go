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
	// reactor := make(map[point]bool)
	reactor := make(map[cuboid]bool)
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

			currCuboid := newCuboid(point{xS, yS, zS}, point{xE, yE, zE}, on)
			cuboids.Push(currCuboid)
			// fmt.Println(currCuboid)
			// fmt.Println(line)
			// fmt.Println(currCuboid.Volume())

			/* for x := xS; x <= xE; x++ {
				for y := yS; y <= yE; y++ {
					for z := zS; z <= zE; z++ {
						fmt.Printf("%d,%d,%d\n", x, y, z)
						reactor[point{x, y, z}] = on

					}
				}
			} */
		} else {
			panic("parsing error")
		}
	}

	// fmt.Println(" ")
	/* for i, c := range cuboids {
	// fmt.Println(c)
	if i == 0 {
		reactor[c] = c.on
		continue
	}
	for j := 0; j < i; j++ {
		// checking if existing cuboids overlap with current one
		fmt.Println(isOverlap(c, cuboids[j]))
		if isOverlap(c, cuboids[j]) {
			Split(c, cuboids[j])
		}

	} */
	// fmt.Printf("There are currently %d 1x1x1 squares on\n", countOnCuboid(reactor))

	/* for _, c1 := range cuboids {
		for _, c2 := range cuboids {
			if isOverlap(c1, c2) {
				// fmt.Println(c1, c2)
			}
		}
	} */

	//try and split first two cuboids
	fmt.Println(cuboids)
	c1, c2 := cuboids.PopLeft(), cuboids.PopLeft()
	fmt.Println(c1, c2)
	fmt.Println(isOverlap(c1, c2))
	Split(c1, c2)
	return countOnCuboid(reactor)
}

func countOn(r map[point]bool) (on int) {
	for _, o := range r {
		if o == true {
			on++
		}
	}
	return
}

// countOnCuboid counts the number of 1 x 1 x 1 cubes that are switched on in the reactor map.
// It assumes that there are no overlapping cuboids, and such overlaps would be handled before insertion into reactor map
func countOnCuboid(r map[cuboid]bool) (on int) {
	for c, o := range r {
		if o {
			on += c.Volume()
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
func Split(c1, c2 cuboid) []cuboid {
	if c1StartInC2(c1, c2) {
		return split(c2, c1)
	} else if c2StartInC1(c1, c2) {
		return split(c1, c2)
	} else {
		return make([]cuboid, 0)
	}
}

func split(a, b cuboid) []cuboid {
	res := make([]cuboid, 0)
	fmt.Println(a, b)
	nc := cuboid{a.start, b.start, true}
	nc2 := cuboid{a.end, b.end, true}
	nc3 := cuboid{b.start, a.end, true}
	nc4 := cuboid{}
	res = append(res, nc, nc2, nc3)
	fmt.Println(res)
	return res
}
