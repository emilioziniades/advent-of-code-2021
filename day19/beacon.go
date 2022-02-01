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

// implementing util.Vector interface
func (p point) ToSlice() []int {
	return []int{p.x, p.y, p.z}
}

type pair struct {
	a, b point
}

func (p pair) String() string {
	return fmt.Sprintf("[(%v),(%v)]", p.a, p.b)
}

type scanner struct {
	points       []point
	id           int
	zeroPosition point // position relative to scanner 0
	neighbours   map[int]neighbour
	done         bool
}

type neighbour struct {
	id             int
	rotationSteps  string
	parentPosition point // position of parent relative to this neighbour
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

// CountBeaconsAndScannerDistance returns the length of the set of beacons (all reoriented to be relative to scanner 0) and the maximum distance between two scanners
func CountBeaconsAndScannerDistance(input []string) (int, int) {

	scanners := parseInputToScanners(input)

	// global set of beacons relative to scanner 0
	allBeacons := make(map[point]bool)

	// adjacency graph between scanners that share 12 points
	adjacency := make(map[int][]int)

	// no reorientation needed, so scanner 0 done
	scanners[0].done = true

	determineSharedBeacons(scanners, adjacency)

	for k, _ := range scanners {
		reorientBeaconsToZero(scanners, k, adjacency)
		reorientScannersToZero(scanners, k, adjacency)
	}

	maxDist := math.MinInt
	posList := make([]point, 0)

	for _, v := range scanners {
		posList = append(posList, v.zeroPosition)
		if !v.done {
			continue
		}
		for _, pt := range v.points {
			allBeacons[pt] = true
		}
	}

	for _, sA := range posList {
		for _, sB := range posList {
			if sA == sB {
				continue
			}
			dist := util.ManhattanDistance(sA, sB)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return len(allBeacons), maxDist
}

/*
Overall, these are the steps to deterrmine shared beacons:

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
	- once you have a direct mapping of overlapping beacons, you can calculate the relative position of scanner A relative to scanner B
	- to determine the relative position, you need to try all possible orientations until scanner A and B and oriented the same way
*/
func determineSharedBeacons(scanners map[int]*scanner, adjacencyGraph map[int][]int) {

	fmt.Println("processing all scanner pairs...")
	n := len(scanners)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			determineSharedBeaconsPairs(scanners[i], scanners[j], adjacencyGraph)
		}
	}
}

func determineSharedBeaconsPairs(scannerA, scannerB *scanner, adjacencyGraph map[int][]int) {

	pairsA := scannerA.makePairs()
	pairsB := scannerB.makePairs()

	// keys are points from scanner A's perspective
	AtoBMap := make(map[point]point)

	sA := stack.New[pair]()
	sB := stack.New[pair]()

	for _, pA := range pairsA {
		for _, pB := range pairsB {
			if distance(pA) == distance(pB) {
				sA.PushLeft(pA)
				sB.PushLeft(pB)
			}
		}
	}

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
				AtoBMap[c.a] = e.a
				AtoBMap[c.b] = e.b
				AtoBMap[d.b] = f.b
			case e.a == f.b:
				AtoBMap[c.a] = e.a
				AtoBMap[c.b] = e.b
				AtoBMap[d.b] = f.a
			case e.b == f.a:
				AtoBMap[c.a] = e.b
				AtoBMap[c.b] = e.a
				AtoBMap[d.b] = f.b
			case e.b == f.b:
				AtoBMap[c.a] = e.b
				AtoBMap[c.b] = e.a
				AtoBMap[d.b] = f.a
			}
		}
	}

	overlapA := make([]point, 0)
	overlapB := make([]point, 0)

	if len(AtoBMap) < 12 {
		// insufficient shared beacons
		return
	}

	for k, v := range AtoBMap {
		overlapA = append(overlapA, k)
		overlapB = append(overlapB, v)
	}

	// this is the position of scannerB relative to scannerA
	var positionBA *point
loop:
	// keep rotating B's beacons until they are in the same orientation as A
	// (we'll know this since the difference between A and B will be a single point - the position of B relative to A's position )
	for index, points := range possibleOrientations(overlapB) {
		for i, pA := range overlapA {
			pB := points[i]
			attempt := point{pA.x - pB.x, pA.y - pB.y, pA.z - pB.z}
			if positionBA == nil {
				positionBA = &attempt
				continue
			}
			if *positionBA == attempt {
				if i == 11 {
					currNeighbour := neighbour{
						parentPosition: *positionBA,
						rotationSteps:  rollTurnDict[index],
						id:             scannerA.id,
					}

					scannerB.neighbours[currNeighbour.id] = currNeighbour
					adjacencyGraph[scannerA.id] = append(adjacencyGraph[scannerA.id], scannerB.id)
					break loop
				}
			} else {
				positionBA = nil
				break
			}
		}
	}
}

// pathToZero finds a possible sequence of reorientations from scanner n to scanner 0 using BFS
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

func reorientBeaconsToZero(scanners map[int]*scanner, id int, adjacencyGraph map[int][]int) {
	path := pathToZero(adjacencyGraph, id)
	// pop starting point off stack
	var curr int
	next := path.PopLeft()
	thisScanner := scanners[id]
	var nextNeighbour neighbour
	for next != 0 {
		curr = next
		next = path.PopLeft()
		nextNeighbour = scanners[curr].neighbours[next]
		// reorient and shift beacon positions
		for i, pt := range thisScanner.points {
			thisScanner.points[i] = shift(rollAndTurn(pt, nextNeighbour.rotationSteps), nextNeighbour.parentPosition)
		}
	}
	thisScanner.done = true
}

func reorientScannersToZero(scanners map[int]*scanner, id int, adjacencyGraph map[int][]int) {
	path := pathToZero(adjacencyGraph, id)
	thisScanner := scanners[id]

	if len(path) <= 1 {
		// scanner 0 already has position (0, 0, 0)
		return
	}

	_ = path.PopLeft()
	curr := path.PopLeft()
	var next int
	thisScanner.zeroPosition = thisScanner.neighbours[curr].parentPosition

	for len(path) > 0 {
		next = path.PopLeft()
		thisScanner.zeroPosition = shift(rollAndTurn(thisScanner.zeroPosition, scanners[curr].neighbours[next].rotationSteps), scanners[curr].neighbours[next].parentPosition)
		curr = next
	}
	return
}

func parseInputToScanners(input []string) map[int]*scanner {
	re := regexp.MustCompile(`---\sscanner\s(\d+)\s---`)
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

func shift(a, b point) point {
	return point{a.x + b.x, a.y + b.y, a.z + b.z}
}
