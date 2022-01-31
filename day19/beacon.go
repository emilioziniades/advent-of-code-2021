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
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

type pair struct {
	a, b point
}

func (p pair) String() string {
	return fmt.Sprintf("[(%v),(%v)]", p.a, p.b)
}

type scanner struct {
	points     []point
	id         int
	neighbours map[int]neighbour
	done       bool
}

type neighbour struct {
	rotationSteps string
	position      point
	id            int
}

func (s scanner) makePairs() []pair {
	pairs := make([]pair, 0)

	for _, e := range s.points {
		for _, f := range s.points {
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
		pairsA: a, b, c
		pairsB: d, e, f
		and the distance(a,b) == distance(d,e), distance(a, c) == distance(d, f)
		then point a must be the same as point d, a == d
		and b == e, and c == f
		and you have the mapping:
		a: d
		b: e
		c: f
	- once you have a direct mapping of overlapping beacons, you can calculate the relative distance of scanner 1 relative to scanner 0
*/

// CountBeacons returns the length of the set of beacons (all reoriented to be relative to scanner 0)
func CountBeacons(input []string) int {

//	test := point{-393,719,612}
//	fmt.Println(unrollAndUnturn("RTTTRTTTRTTTRTRRTTTRTTTR", test))
//	return 0

	scanners := parseInputToScanners(input)

	// global set of beacons relative to scanner 0
	allBeacons := make(map[point]bool)
	adjacency := make(map[int][]int)

	// no reorientation needed, so scanner 0 done
	scanners[0].done = true

	determineSharedBeacons(allBeacons, scanners, adjacency)

	for k, _ := range scanners {
		if k < 2 || k == 4 {
			reorientToZero(scanners, k, adjacency)
		}
	}
	for _, v := range scanners {
		curr := *v
		fmt.Printf("scanner id: %d, done: %v\n", curr.id, curr.done)
		for _, e := range curr.neighbours {
			fmt.Printf("\tneighbour: %d, relative position: %v, steps: %s\n", e.id, e.position, e.rotationSteps)
		}
	}

	for _, v := range scanners {
		c := 0
		if !v.done {
			continue
		}
		for _, pt := range v.points {
			_, ok := allBeacons[pt]
			if !ok {
				c++
			}
			allBeacons[pt] = true
		}
		fmt.Printf("scanner %d added %d new beacons\n", v.id, c)
	}
	return len(allBeacons)
}

func determineSharedBeacons(beaconSet map[point]bool, scanners map[int]*scanner, adjacencyGraph map[int][]int) {

	n := len(scanners)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			fmt.Printf("----------Processing scanner %d and scanner %d----------\n", j, i)
			determineSharedBeaconsPairs(beaconSet, scanners[j], scanners[i], adjacencyGraph)
		}
	}

	fmt.Println(adjacencyGraph)
	for k, _ := range scanners {
		fmt.Println(k, " : ", pathToZero(adjacencyGraph, k))
	}
}

func determineSharedBeaconsPairs(beaconSet map[point]bool, scannerA, scannerB *scanner, adjacencyGraph map[int][]int) {

	pairsA := scannerA.makePairs()
	pairsB := scannerB.makePairs()

	// keys are points from scanner 0's perspective
	bm := make(map[point]point)

	sA := stack.New[pair]()
	sB := stack.New[pair]()

	count := 0

	for _, pA := range pairsA {
		for _, pB := range pairsB {
			if d0, d1 := distance(pA), distance(pB); d0 == d1 {
				count++
				//				fmt.Printf("equal distance = %f between points %v and %v\n", d0, pA, pB)
				sA.PushLeft(pA)
				sB.PushLeft(pB)
			}
		}
	}
	fmt.Printf("\tFound %d pairs with equal distances\n", count)

	for len(sA) > 0 && len(sB) > 0 {

		if len(sA) == 1 && len(sB) == 1 {
			break
		}

		c := sA.Pop()
		d := sA.Pop()

		e := sB.Pop()
		f := sB.Pop()

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

	overlapA := make([]point, 0)
	overlapB := make([]point, 0)

	fmt.Printf("\t%d common beacons determined\n", len(bm))
	if len(bm) < 12 {
		fmt.Println("\tnot enough common beacons")
		return
	}

	for k, v := range bm {
		fmt.Println(k, " :: ", v)
		overlapA = append(overlapA, k)
		overlapB = append(overlapB, v)
	}

	// this is the position of scannerB relative to scannerA
	var positionBA *point
loop:
	for index, points := range possibleOrientations(overlapB) {
		//		fmt.Println(points)
		for i, pA := range overlapA {
			pB := points[i]
			attempt := point{pA.x - pB.x, pA.y - pB.y, pA.z - pB.z}
			if positionBA == nil {
				positionBA = &attempt
				continue
			}
			if *positionBA == attempt {
				//					fmt.Println(pA, pB, attempt)
				//					fmt.Println("works so far")
				if i == 11 {
					//TODO currently, this finds position relative to scannerA, and more work required to find position relative to scanner0, which may not necessarily be scannerA
					fmt.Println("\tFound scanner position!!! - ", positionBA)
					currNeighbour := neighbour{
						position:      *positionBA,
						rotationSteps: rollTurnDict[index],
						id:            scannerA.id,
					}

					scannerB.neighbours[currNeighbour.id] = currNeighbour
					adjacencyGraph[scannerA.id] = append(adjacencyGraph[scannerA.id], scannerB.id)
					break loop
				}
			} else {
				//fmt.Printf("doesn't work. distance between %v and %v is %v, but want %v\n", pA, pB, attempt, positionBA)
				positionBA = nil
				break
			}
		}
	}
	/*
			// do the reverse rotation for all scannerB beacons
			fmt.Printf("Currently reorienting scanners from scanner %d to be relative to %d\n", scannerB.id, scannerB.relativeTo)
			for i, e := range scannerB.points {
			// do reverse of current orientation for all of scannerB's points,
		// so that they are in the correct orientation relative to scanner A
				reoriented := unrollAndUnturn(scannerB.rotationSteps, e)
				shifted := point{reoriented.x + scannerB.position.x, reoriented.y + scannerB.position.y, reoriented.z + scannerB.position.z}
				scannerB.points[i] = shifted
				if scannerB.relativeTo == 0 {
					beaconSet[shifted] = true
				}

			}
			fmt.Printf("scanner %d positions now relative to scanner %d\n", scannerB.id, scannerB.relativeTo)
		//	*/
}

// pathToZero finds a possible sequence of reorientations from scanner n to scanner 0
func pathToZero(adjacencyGraph map[int][]int, start int) stack.Stack[int] {
	end := 0
	q := stack.New[int]()
	q.PushLeft(start)
	cameFrom := make(map[int]int)
	cameFrom[start] = -1

	for len(q) > 0 {
		curr := q.Pop()
		for _, next := range adjacencyGraph[curr] {
			if _, ok := cameFrom[next]; !ok {
				q.PushLeft(next)
				cameFrom[next] = curr
			}
		}

	}
	path := stack.New[int]()

	curr := end
	for curr != start {
		path.Push(curr)
		curr = cameFrom[curr]
	}
	path.Push(start)
	util.Reverse(path)
	return path
}

func reorientToZero(scanners map[int]*scanner, id int, adjacencyGraph map[int][]int) {
	path := pathToZero(adjacencyGraph, id)
	fmt.Printf("----------Reorienting points from scanner %d to be relative to scanner 0 via path %v----------\n", id, path)

	// pop starting point off stack
	var prev int
	curr := path.PopLeft()
	thisScanner := scanners[id]
	var currNeighbour neighbour
	for curr != 0 {
		prev = curr
		curr = path.PopLeft()
		currNeighbour = scanners[prev].neighbours[curr]
		fmt.Printf("\treorienting from scanner %d --> scanner %d\n", prev, curr)
		fmt.Printf("\tUsing rotation steps %s and relative position %v of scanner %d from perspective of %d\n", currNeighbour.rotationSteps, currNeighbour.position, currNeighbour.id, prev)
		fmt.Println("\t", currNeighbour)

		// reorient and shift scanner position
		//		reorientScanner := unrollAndUnturn(currNeighbour.rotationSteps, thisScanner.position)
		//		shiftScanner :=

		// reorient and shift beacon positions
		for i, pt := range thisScanner.points {
			reorientedPt := unrollAndUnturn(currNeighbour.rotationSteps, pt)
			shiftedPt := point{reorientedPt.x + currNeighbour.position.x, reorientedPt.y + currNeighbour.position.y, reorientedPt.z + currNeighbour.position.z}
			thisScanner.points[i] = shiftedPt

		}
	}
//	fmt.Println("\t\t", thisScanner.points)
	thisScanner.done = true
}

func parseInputToScanners(input []string) map[int]*scanner {
	re := regexp.MustCompile(`---\sscanner\s(\d)\s---`)
	var currentScanner int
	scanners := make(map[int]*scanner)

	for _, in := range input {
		if in == "" {
			continue
		}
		if re.MatchString(in) {
			match := re.FindStringSubmatch(in)
			num, _ := strconv.Atoi(match[1])
			currentScanner = num
			scanners[currentScanner] = &scanner{id: currentScanner, neighbours: make(map[int]neighbour)}
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
			if s, ok := scanners[currentScanner]; ok {
				s.points = append(s.points, point{x: nums[0], y: nums[1], z: nums[2]})
				scanners[currentScanner] = s
			}
		}
	}

	return scanners
}

func distance(p pair) float64 {
	diffx := float64(p.a.x - p.b.x)
	diffy := float64(p.a.y - p.b.y)
	diffz := float64(p.a.z - p.b.z)
	return math.Sqrt(math.Pow(diffx, 2) + math.Pow(diffy, 2) + math.Pow(diffz, 2))
}
