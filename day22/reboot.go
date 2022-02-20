package day22

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/emilioziniades/adventofcode2021/stack"
)

type Point struct {
	x, y, z int
}

type Cuboid struct {
	start, end Point
	on         bool
}

func NewCuboid(a, b Point, on bool) *Cuboid {
	/* format := "NewCuboid: start %s greater than end %s. %v %v"
	if a.x > b.x {
		log.Fatalf(format, "x", "x", a, b)
	}
	if a.y > b.y {
		log.Fatalf(format, "y", "y", a, b)
	}
	if a.z > b.z {
		log.Fatalf(format, "z", "z", a, b)
	} */
	return &Cuboid{a, b, on}
}

func (c Cuboid) Volume() int {
	return (c.end.x - c.start.x + 1) * (c.end.y - c.start.y + 1) * (c.end.z - c.start.z + 1)
}

type toggler struct {
	x1, x2, toggleVal int
}

func newToggler(x1, x2 int) toggler {
	return toggler{
		toggleVal: x1 ^ x2,
		x1:        x1,
		x2:        x2,
	}
}

func (t toggler) toggle(x int) int {
	if x != t.x1 && x != t.x2 {
		panic("toggle: can't toggle value not included in toggleVal")
	}
	return t.toggleVal ^ x
}

func Reboot(input []string, bounds float64) int {
	reactor := make(map[Cuboid]bool)
	Cuboids := inputToCuboids(input, bounds)
	fmt.Printf("%#v", Cuboids)

	/*
		The logic for adding cubes to the reactor map is as follows:
		For every Cuboid from the input:
			If not intersected with any other existing Cuboid in reactor:
				Add to reactor
			Else if there is intersection:
				Split two intersecting Cuboids into up to 7 child Cuboids, with appropriate on/off status
				Then check if those children Cuboids intersect with existing Cuboids in reactor:
					If no intersect:
						Add to reactor
					If intersect:
						Recursively do the above procedure

	*/

	var recReboot func(Cuboid)
	recReboot = func(c Cuboid) {
		fmt.Printf("checking Cuboid %v\n", c)
		for r := range reactor {
			//check for overlap
			if isOverlap(c, r) {
				fmt.Printf("found overlap between %v and %v\n", c, r)
				children := Split(c, r)
				//delete overlapping Cuboid since its children will replace it
				delete(reactor, r)
				for _, child := range children {
					recReboot(child)
				}
				return
			}
		}
		// if got here, no overlap, add to reactor
		fmt.Printf("inserting Cuboid %v into reactor\n", c)
		reactor[c] = c.on
	}

	for len(Cuboids) > 0 {
		// pop next Cuboid
		curr := Cuboids.PopLeft()
		fmt.Printf("********* next instruction: %v *********\n", curr)
		recReboot(curr)
		fmt.Printf("********* current unit cubes count: %d *********\n", countOnCuboid(reactor))
	}
	return countOnCuboid(reactor)
}

// when two Cuboids overlap, they can be split into seven distinct Cuboids and treated separately
func Split(s1, s2 Cuboid) (children []Cuboid) {
	/*
		This is looking with positive x right, positive y away, and positive z up

		1. Cuboid to the left before intersection
		2. Cuboid on near side before intersection
		3. Cuboid underneath before intersection (always directly below 4)
		4. Cuboid which overlaps with both
		5. Cuboid on top after intersection (always directly above 4)
		6. Cuboid on far side after intersection
		7. Cuboid to the right after intersection

		Parent Split function will always provide a and b such that b's starting point is within a
		But one must also check if a's end point is also within b
	*/

	xS := []int{s1.start.x, s1.end.x, s2.start.x, s2.end.x}
	yS := []int{s1.start.y, s1.end.y, s2.start.y, s2.end.y}
	zS := []int{s1.start.z, s1.end.z, s2.start.z, s2.end.z}
	sort.Ints(xS)
	sort.Ints(yS)
	sort.Ints(zS)

	x1, x2, x3, x4 := xS[0], xS[1], xS[2], xS[3]
	y1, y2, y3, y4 := yS[0], yS[1], yS[2], yS[3]
	z1, z2, z3, z4 := zS[0], zS[1], zS[2], zS[3]

	var c1, c2, c3, c4, c5, c6, c7 *Cuboid
	res := make([]*Cuboid, 0)

	// default state is s1 wholly contains s2 on all sides
	//TODO determine if children cuboids are on or off
	c1 = &Cuboid{Point{x1, y1, z1}, Point{x2 - 1, y4, z4}, s1.on}
	c2 = &Cuboid{Point{x2, y1, z1}, Point{x3, y4, z2 - 1}, s1.on}
	c3 = &Cuboid{Point{x2, y1, z2}, Point{x3, y2 - 1, z3}, s1.on}
	c4 = &Cuboid{Point{x2, y2, z2}, Point{x3, y3, z3}, s2.on}
	c5 = &Cuboid{Point{x2, y3 + 1, z2}, Point{x3, y4, z3}, s1.on}
	c6 = &Cuboid{Point{x2, y1, z3 + 1}, Point{x3, y4, z4}, s1.on}
	c7 = &Cuboid{Point{x3 + 1, y1, z1}, Point{x4, y4, z4}, s1.on}

	y1y2 := newToggler(y1, y2)
	y3y4 := newToggler(y3, y4)
	z1z2 := newToggler(z1, z2)
	z3z4 := newToggler(z3, z4)

	if s1.start.x > s2.start.x {
		c1.start.y = y1y2.toggle(c1.start.y)
		c1.start.z = z1z2.toggle(c1.start.z)
		c1.end.y = y3y4.toggle(c1.end.y)
		c1.end.z = z3z4.toggle(c1.end.z)

		c1.on = s2.on
	}

	if s1.start.y > s2.start.y {
		c1.start.y = y1y2.toggle(c1.start.y)
		c2.start.y = y1y2.toggle(c2.start.y)
		c6.start.y = y1y2.toggle(c6.start.y)
		c7.start.y = y1y2.toggle(c7.start.y)

		c3.on = s2.on

	}

	if s1.start.z > s2.start.z {
		c1.start.z = z1z2.toggle(c1.start.z)
		c2.start.y = y1y2.toggle(c2.start.y)
		c2.end.y = y3y4.toggle(c2.end.y)
		c7.start.z = z1z2.toggle(c7.start.z)

		c2.on = s2.on
	}

	if s1.end.x < s2.end.x {
		c7.start.y = y1y2.toggle(c7.start.y)
		c7.start.z = z1z2.toggle(c7.start.z)
		c7.end.y = y3y4.toggle(c7.end.y)
		c7.end.z = z3z4.toggle(c7.end.z)

		c7.on = s2.on
	}

	if s1.end.y < s2.end.y {
		c1.end.y = y3y4.toggle(c1.end.y)
		c2.end.y = y3y4.toggle(c2.end.y)
		c6.end.y = y3y4.toggle(c6.end.y)
		c7.end.y = y3y4.toggle(c7.end.y)

		c5.on = s2.on
	}

	if s1.end.z < s2.end.z {
		c1.end.z = z3z4.toggle(c1.end.z)
		c6.start.y = y1y2.toggle(c6.start.y)
		c6.end.y = y3y4.toggle(c6.end.y)
		c7.end.z = z3z4.toggle(c7.end.z)

		c6.on = s2.on
	}

	if s1.start.x == s2.start.x {
		c1 = nil
	}
	if s1.start.z == s2.start.z {
		c2 = nil
	}
	if s1.start.y == s2.start.y {
		c3 = nil
	}
	if s1.end.y == s2.end.y {
		c5 = nil
	}
	if s1.end.z == s2.end.z {
		c6 = nil
	}
	if s1.end.x == s2.end.x {
		c7 = nil
	}

	res = append(res, c1, c2, c3, c4, c5, c6, c7)

	for _, c := range res {
		if c != nil {
			children = append(children, *c)
		}
	}
	return children

}

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

// countOnCuboid counts the number of unit cubes that are on in the reactor map. It assumes that there are no overlapping Cuboids, and such overlaps would be handled before insertion into reactor map
func countOnCuboid(r map[Cuboid]bool) (on int) {
	for c, o := range r {
		if o {
			on += c.Volume()
		}
	}
	return
}

func inputToCuboids(input []string, bounds float64) stack.Stack[Cuboid] {
	Cuboids := stack.New[Cuboid]()
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

			currCuboid := Cuboid{Point{xS, yS, zS}, Point{xE, yE, zE}, on}
			Cuboids.Push(currCuboid)

		} else {
			panic("parsing error")
		}
	}
	return Cuboids
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
