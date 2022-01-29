package day19

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/emilioziniades/adventofcode2021/stack"
	"github.com/emilioziniades/adventofcode2021/util"
)

type point struct {
	x, y, z int
}

func (p point) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

type pair struct {
	a, b point
}

func (p pair) String() string {
	return fmt.Sprintf("[(%v),(%v)]", p.a, p.b)
}

type report []point

func (r report) makePairs() []pair {
	pairs := make([]pair, 0)

	for _, e := range r {
		for _, f := range r {
			if e == f {
				continue
			}
			//since distance is symmetrical, only need pair{a,b}, not also pair{b,a}
			if !util.Has(pairs, pair{f, e}) {
				curr := pair{e, f}
				pairs = append(pairs, curr)
			}
		}
	}
	return pairs
}

/*
Overall, these are the steps:

1. For each pair of scanners, do the following:
	- consider all the pairs, and then compare distances between reports to find pairs which have the same distance
	- using this, try and isolate each beacon. If you have two pairs like this:
		scanner0: a, b, c
		scanner1: d, e, f
		and the distance(a,b) == distance(d,e), distance(a, c) == distance(d, f)
		then point a must be the same as point d, a == d
		and b == e, and c == f
		and you have the mapping:
		a: d
		b: e
		c: f
	- once you have a direct mapping of overlapping beacons, you can calculate the relative distance of scanner 1 relative to scanner 0
*/

func CountBeacons(input []string) int {

	reports := parseInputToReports(input)

	// global set of beacons relative to scanner 0
	// countBeacons returns the length of this set
	allBeacons := make(map[point]bool)

	// first populate with beacons from scanner 0 report
	for _, e := range reports[0] {
		allBeacons[e] = true
	}

	n := len(reports)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				break
			}
			determineSharedBeacons(&allBeacons, reports[i], reports[j])
		}
	}

	scanner0 := reports[0].makePairs()
	scanner1 := reports[1].makePairs()

	// keys are points from scanner 0's perspective
	bm := make(map[point]point)

	s0 := stack.New[pair]()
	s1 := stack.New[pair]()

	count := 0

	for _, p0 := range scanner0 {
		for _, p1 := range scanner1 {
			if d0, d1 := distance(p0), distance(p1); d0 == d1 {
				count++
				fmt.Printf("equal distance = %f between points %v and %v\n", d0, p0, p1)
				s0.PushLeft(p0)
				s1.PushLeft(p1)
			}
		}
	}
	fmt.Println(count)

	for len(s0) > 0 && len(s1) > 0 {
		c := s0.Pop()
		d := s0.Pop()

		e := s1.Pop()
		f := s1.Pop()

		fmt.Println(c, d, e, f)

		if c.a == d.a {
			switch {
			case e.a == f.a:
				bm[c.a] = e.a
				bm[c.b] = e.b
				bm[d.b] = f.b
			case e.a == f.b:
				bm[c.a] = e.a
				bm[c.b] = e.b
				bm[d.b] = f.a
			case e.b == f.a:
				bm[c.a] = e.b
				bm[c.b] = e.a
				bm[d.b] = f.b
			case e.b == f.b:
				bm[c.a] = e.b
				bm[c.b] = e.a
				bm[d.b] = f.a
			}
		}
	}

	overlap0 := make([]point, 0)
	overlap1 := make([]point, 0)

	for k, v := range bm {
		fmt.Println(k, " :: ", v)
		overlap0 = append(overlap0, k)
		overlap1 = append(overlap1, v)
	}

	steps := ""
	newOverlap1 := make([]point, 0)
	var position1 *point
loop:
	for index, points := range possibleOrientations(overlap1) {
		//		fmt.Println(points)
		for i, e0 := range overlap0 {
			e1 := points[i]
			attempt := point{e0.x - e1.x, e0.y - e1.y, e0.z - e1.z}
			if position1 == nil {
				position1 = &attempt
				continue
			}
			if *position1 == attempt {
				//					fmt.Println(e0, e1, attempt)
				//					fmt.Println("works so far")
				if i == 11 {
					fmt.Println("Found scanner position!!! - ", position1)
					// ensures that we have scanner 1's beacons in the correct orientation relative to scanner 0
					newOverlap1 = points
					steps = rollTurnDict[index]
					fmt.Println(steps)
					break loop
				}
			} else {
				//					fmt.Printf("doesn't work. distance between %v and %v is %v, but want %v\n", e0, e1, attempt, position1)
				position1 = nil
				break
			}
		}
	}

	fmt.Println(overlap1)
	fmt.Println(newOverlap1)

	// do the reverse rotation for all scanner 1 beacons
	for _, e := range reports[1] {
		reoriented := unrollAndUnturn(steps, e)
		shifted := point{reoriented.x + position1.x, reoriented.y + position1.y, reoriented.z + position1.z}
		fmt.Printf("%v / %v / %v \n", e, reoriented, shifted)
		allBeacons[shifted] = true
	}

	return len(allBeacons)
}

func determineSharedBeacons(beaconSet *map[point]bool, reportA, reportB report) {

}

func parseInputToReports(input []string) map[int]report {
	re := regexp.MustCompile(`---\sscanner\s(\d)\s---`)
	var currentScanner int
	reports := make(map[int]report)

	for _, in := range input {
		if in == "" {
			continue
		}
		if re.MatchString(in) {
			match := re.FindStringSubmatch(in)
			num, _ := strconv.Atoi(match[1])
			currentScanner = num
		} else {
			numsStr := strings.Split(in, ",")
			nums := make([]int, 0)
			for _, e := range numsStr {
				num, err := strconv.Atoi(e)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, num)
			}
			reports[currentScanner] = append(reports[currentScanner], point{x: nums[0], y: nums[1], z: nums[2]})
		}
	}
	return reports

}

func distance(p pair) float64 {
	diffx := float64(p.a.x - p.b.x)
	diffy := float64(p.a.y - p.b.y)
	diffz := float64(p.a.z - p.b.z)
	return math.Sqrt(math.Pow(diffx, 2) + math.Pow(diffy, 2) + math.Pow(diffz, 2))
}
