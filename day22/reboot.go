package day22

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/emilioziniades/adventofcode2021/stack"
)

type Point struct {
	X, Y, Z int
}

type Cuboid struct {
	Start, End Point
	On         bool
}

func NewCuboid(a, b Point, on bool) *Cuboid {
	/* format := "NewCuboid: start %s greater than end %s. %v %v"
	if a.X > b.X {
		log.Fatalf(format, "X", "X", a, b)
	}
	if a.Y > b.Y {
		log.Fatalf(format, "Y", "Y", a, b)
	}
	if a.Z > b.Z {
		log.Fatalf(format, "Z", "Z", a, b)
	} */
	return &Cuboid{a, b, on}
}

func (c Cuboid) Volume() int {
	return (c.End.X - c.Start.X + 1) * (c.End.Y - c.Start.Y + 1) * (c.End.Z - c.Start.Z + 1)
}

func Reboot(input []string, bounds float64) int {
	reactor := make(map[Cuboid]bool)
	Cuboids := inputToCuboids(input, bounds)
	fmt.Println(Cuboids)

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
		reactor[c] = c.On
	}

	for len(Cuboids) > 1 {
		// pop next Cuboid
		curr := Cuboids.PopLeft()
		fmt.Printf("********* next instruction: %v *********\n", curr)
		recReboot(curr)
		fmt.Printf("********* current unit cubes count: %d *********\n", countOnCuboid(reactor))
	}
	return countOnCuboid(reactor)
}

func isOverlap(a, b Cuboid) bool {
	//does one cube have a start or end Point that lies within the other cube?
	return isAStartInB(a, b) || isAStartInB(b, a)
}

func isAStartInB(a, b Cuboid) bool {
	return (a.Start.X >= b.Start.X && a.Start.X <= b.End.X) && (a.Start.Y >= b.Start.Y && a.Start.Y <= b.End.Y) && (a.Start.Z >= b.Start.Z && a.Start.Z <= b.End.Z)

}

func isAEndInB(a, b Cuboid) bool {
	return (a.End.X >= b.Start.X && a.End.X <= b.End.X) && (a.End.Y >= b.Start.Y && a.End.Y <= b.End.Y) && (a.End.Z >= b.Start.Z && a.End.Z <= b.End.Z)
}

// when two Cuboids overlap, they can be split into seven distinct Cuboids and treated separately
func Split(c1, c2 Cuboid) []Cuboid {
	if isAStartInB(c1, c2) {
		return split(c2, c1, true)
	} else if isAStartInB(c2, c1) {
		return split(c1, c2, false)
	} else {
		return make([]Cuboid, 0)
	}
}

// for now, assuming that one Cuboid does not contain the other Cuboid
func split(a, b Cuboid, inverted bool) []Cuboid {
	res := make([]*Cuboid, 0)
	fmt.Println("a: ", a, a.Volume())
	fmt.Println("b: ", b, b.Volume())

	// overlapping section
	var overlapOn bool
	if inverted {
		overlapOn = a.On
	} else {
		overlapOn = b.On
	}
	/*
		This is looking with positive X right, positive Y away, and positive Z up

		1. Cuboid to the left before intersection
		2. Cuboid to the right after intersection
		3. Cuboid which overlaps with both
		4. Cuboid on near side before intersection
		5. Cuboid underneath before intersection (always directly below 3)
		6. Cuboid on far side after intersection
		7. Cuboid on top after intersection (always directly above 3)

		Parent Split function will always provide a and b such that b's starting point is within a
		But one must also check if a's end point is also within b
	*/

	var nc1, nc2, nc3, nc4, nc5, nc6, nc7 *Cuboid
	if isAEndInB(a, b) {
		// original splits
		nc1 = NewCuboid(a.Start, Point{b.Start.X - 1, a.End.Y, a.End.Z}, a.On) //same in both case
		nc2 = NewCuboid(Point{a.End.X + 1, b.Start.Y, b.Start.Z}, b.End, b.On)
		nc3 = NewCuboid(b.Start, a.End, overlapOn)
		nc4 = NewCuboid(Point{b.Start.X, a.Start.Y, a.Start.Z}, Point{a.End.X, b.Start.Y - 1, a.End.Z}, a.On)
		nc5 = NewCuboid(Point{b.Start.X, b.Start.Y, a.Start.Z}, Point{a.End.X, a.End.Y, b.Start.Z - 1}, a.On)
		nc6 = NewCuboid(Point{b.Start.X, a.End.Y + 1, b.Start.Z}, Point{a.End.X, b.End.Y, b.End.Z}, b.On)
		nc7 = NewCuboid(Point{b.Start.X, b.Start.Y, a.End.Z + 1}, Point{a.End.X, a.End.Y, b.End.Z}, b.On)
	} else {
		nc1 = NewCuboid(a.Start, Point{b.Start.X - 1, a.End.Y, a.End.Z}, a.On) //same in both cases
		nc2 = NewCuboid(Point{b.End.X + 1, a.Start.Y, a.Start.Z}, a.End, b.On)
		nc3 = NewCuboid(b.Start, Point{b.End.X, a.End.Y, a.End.Z}, overlapOn)
		nc4 = NewCuboid(Point{b.Start.X, a.Start.Y, a.Start.Z}, Point{b.End.X, b.Start.Y - 1, a.End.Z}, a.On)
		nc5 = NewCuboid(Point{b.Start.X, b.Start.Y, a.Start.Z}, Point{b.End.X, a.End.Y, b.Start.Z - 1}, a.On)
		nc6 = NewCuboid(Point{b.Start.X, a.End.Y + 1, b.Start.Z}, Point{b.End.X, b.End.Y, b.End.Z}, b.On)
		nc7 = NewCuboid(Point{b.Start.X, b.Start.Y, a.End.Z + 1}, Point{b.End.X, a.End.Y, b.End.Z}, b.On)
	}
	if a.Start.X == b.Start.X {
		//left side shared, no cube 1
		nc1 = nil
	}
	if a.End.X == b.End.X {
		// right side shared, no cube 2
		nc2 = nil
	}

	res = append(res, nc1, nc2, nc3, nc4, nc5, nc6, nc7)
	final := make([]Cuboid, 0)
	sum := 0
	fmt.Println("children: ")
	for i, c := range res {
		if c == nil {
			continue
		}
		if c.Volume() <= 0 {
			fmt.Println("something whacky going on down below \\/")
		}
		fmt.Println(i+1, *c, c.Volume())
		sum += c.Volume()
		final = append(final, *c)
	}
	fmt.Println(sum)
	return final
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
