package day23

import (
	"fmt"
	"reflect"

	"github.com/emilioziniades/adventofcode2021/queue"
)

const (
	hallRow      = 1
	upperHomeRow = 2 // row closer to hall
	lowerHomeRow = 3 // row furthest from hall

	AHomeCol     = 3
	BHomeCol     = 5
	CHomeCol     = 7
	DHomeCol     = 9
	hallStartCol = 1
	hallEndCol   = 11
)

var (
	homePositions = map[rune][2]Point{
		'A': {{2, 3}, {3, 3}},
		'B': {{2, 5}, {3, 5}},
		'C': {{2, 7}, {3, 7}},
		'D': {{2, 9}, {3, 9}},
	}

	homeColumn = map[rune]int{
		'A': 3,
		'B': 5,
		'C': 7,
		'D': 9,
	}

	travelCost = map[rune]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}

	endState = State{
		{upperHomeRow, AHomeCol}: 'A',
		{lowerHomeRow, AHomeCol}: 'A',
		{upperHomeRow, BHomeCol}: 'B',
		{lowerHomeRow, BHomeCol}: 'B',
		{upperHomeRow, CHomeCol}: 'C',
		{lowerHomeRow, CHomeCol}: 'C',
		{upperHomeRow, DHomeCol}: 'D',
		{lowerHomeRow, DHomeCol}: 'D',
	}
)

func Djikstra(file string) {
	startState := ParseInitialState(file)

	frontier := queue.NewPriority[*State]()
	frontier.Enqueue(&startState, 0)
	cameFrom := make(map[*State]*State)
	costSoFar := make(map[*State]int)
	cameFrom[&startState] = nil
	costSoFar[&startState] = 0

	for len(frontier) != 0 {

		current := frontier.Dequeue().Value
		PrintState(current)

		if reflect.DeepEqual(current, endState) {
			fmt.Println("found end state!")
			break
		}

		for next, cost := range getStateNeighbours(current) {
			newCost := costSoFar[current] + cost

			if cost, ok := costSoFar[next]; !ok || newCost < cost {
				costSoFar[next] = newCost
				frontier.Enqueue(next, newCost)
				cameFrom[next] = current
			}
		}

	}

	// pp.Println(costSoFar)

}

func getStateNeighbours(currentState *State) map[*State]int {
	stateNeighbours := make(map[*State]int)

	// next possible states are all the possible ways each amphipod can move
	for podPosition, podType := range *currentState {
		if podPosition.IsHome(podType) {
			continue
		}

		// get all possible next positions for this pod, and add those to state neighbours
		nextPodPositions := GetPodNextPositionsAndCosts(podPosition, podType, currentState)
		for nextPodPosition, cost := range nextPodPositions {
			nextState := (*currentState).CopyWithout(podPosition)
			(*nextState)[nextPodPosition] = podType
			stateNeighbours[nextState] = cost
		}
	}

	return stateNeighbours
}

func GetPodNextPositionsAndCosts(podPosition Point, podType rune, state *State) map[Point]int {

	if _, ok := (*state)[podPosition]; !ok {
		panic("current pod not in existing state")
	}
	nextPositions := make(map[Point]int, 0)
	homePosition := homePositions[podType]
	stateMap := *state

	if podPosition.IsHome(podType) && !HomeButMustMakeSpace(podPosition, podType, state) {
		// home already, nowhere to go
		return nextPositions
	}

	if podPosition.InHallway() {
		// only place it can go is home

		// can it reach home?

		// which direction should it go?
		goLeft := homePosition[0].Col < podPosition.Col
		var delta int
		if goLeft {
			delta = -1
		} else {
			delta = 1
		}

		// can it get home?
		targetCol := homePosition[0].Col
		for i := podPosition.Col; i != targetCol; i += delta {
			// move in direction of home
			currentPosition := Point{hallRow, i}
			if _, blocked := stateMap[currentPosition]; blocked {
				// blocked, can't get home
				return nextPositions

			}
		}

		// it can get home!

		// try lower slot, then upper slot
		if _, ok := stateMap[homePosition[1]]; !ok {
			// lower position free
			nextPositions[homePosition[1]] = DistanceCost(podPosition, homePosition[1], podType)
			return nextPositions

		} else if _, ok := stateMap[homePosition[0]]; ok {
			// lower not free, upper free
			nextPositions[homePosition[0]] = DistanceCost(podPosition, homePosition[0], podType)
			return nextPositions
		} else {
			panic("amphipod got nowhere to go! both it's homes are used up")
		}
	}

	// it's not home, it's not in hallway, it must be in starting position!

	// can't move if anyone above
	_, podIsAbove := stateMap[Point{upperHomeRow, podPosition.Col}]
	if podPosition.Row == lowerHomeRow && podIsAbove {
		return nextPositions

	}
	if _, blocked := stateMap[homePosition[0]]; podPosition == homePosition[1] && blocked {
		// it's in lower slot and someone above it
		// this pod can't move
		return nextPositions
	}

	// it can get into hallway

	// go as far left as possible
	for col := podPosition.Col; col >= hallStartCol; col-- {
		currentPosition := Point{hallRow, col}

		if col == AHomeCol || col == BHomeCol || col == CHomeCol || col == DHomeCol {
			// above home, cant stop here
			continue
		}
		if _, blocked := stateMap[currentPosition]; blocked {
			break
		}

		nextPositions[currentPosition] = DistanceCost(podPosition, currentPosition, podType)

	}

	// go as far right as possible
	for col := podPosition.Col; col <= hallEndCol; col++ {
		currentPosition := Point{hallRow, col}

		if col == AHomeCol || col == BHomeCol || col == CHomeCol || col == DHomeCol {
			// above home, cant stop here
			continue
		}
		if _, blocked := stateMap[currentPosition]; blocked {
			break
		}

		nextPositions[currentPosition] = DistanceCost(podPosition, currentPosition, podType)
	}
	return nextPositions
}

func HomeButMustMakeSpace(podPosition Point, podType rune, state *State) bool {
	// if
	// - pod is home
	// - pod in upper home row
	// - pod in row below is wrong type
	// the pod is home, but it needs to let the other pod out

	isHome := podPosition.IsHome(podType)
	isInUpperRow := podPosition.Row == upperHomeRow

	if isInUpperRow && isHome {
		lowerPos := Point{lowerHomeRow, podPosition.Col}
		if lowerPod, ok := (*state)[lowerPos]; ok {
			// someone is in the spot below
			if !lowerPos.IsHome(lowerPod) {
				return true
			}

		}
	}

	return false

}

func (s *State) Copy() *State {
	newState := make(State)
	for k, v := range *s {
		newState[k] = v
	}
	return &newState
}

func (s *State) CopyWithout(p Point) *State {
	newState := s.Copy()
	delete(*newState, p)
	return newState
}

func (p Point) IsHome(a rune) bool {
	if p.InHallway() {
		// in hallway
		return false
	}
	if p.Col == homeColumn[a] {
		// not in hallway, and in home column => home
		return true
	}
	return false
}

func (p Point) InHallway() bool {
	return p.Row == hallRow
}

func DistanceCost(src, dst Point, podType rune) int {
	dist := Distance(src, dst)
	return dist * travelCost[podType]
}

func Distance(src, dst Point) int {
	var dist int

	// distance up to corridor
	dist += (src.Row - hallRow)

	// distance along corridor
	dist += Abs(src.Col - dst.Col)

	// distance down from corridor
	dist += (dst.Row - hallRow)

	return dist
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func PrintState(state *State) {
	for row := 1; row <= 3; row++ {
		for col := 1; col <= 11; col++ {
			point := Point{row, col}
			podType, ok := (*state)[point]
			if ok {
				fmt.Print(string(podType))
			} else if row == hallRow {
				fmt.Print(".")
			} else if col%2 != 0 && col != hallEndCol && col != hallStartCol {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
