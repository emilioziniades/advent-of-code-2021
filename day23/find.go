package day23

import (
	"fmt"
	"strings"

	"github.com/emilioziniades/adventofcode2021/queue"
)

type StateNeighbour struct {
	state State
	cost  int
}

const (
	debug = false

	hallRow = 1

	hallStartCol = 1
	hallEndCol   = 11

	smallState = 8
	largeState = 16
)

var (
	homeColumns = map[string]int{
		"A": 3,
		"B": 5,
		"C": 7,
		"D": 9,
	}

	travelCost = map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}
)

func Djikstra(file string) int {
	startState := ParseState(file)
	endState := GetEndState(startState)

	frontier := queue.NewPriority[State]()
	frontier.Enqueue(startState, 0)
	cameFrom := make(map[string]string)
	costSoFar := make(map[string]int)
	cameFrom[startState.String()] = ""
	costSoFar[startState.String()] = 0

	for len(frontier) != 0 {

		currentState := frontier.Dequeue().Value
		if debug {
			fmt.Println("CURRENT: ")
			fmt.Println()
			PrintState(currentState)
			fmt.Println()
		}

		if currentState.String() == endState.String() {
			if debug {
				fmt.Println("found end state!")
				PrintPath(startState, cameFrom)
			}
			break
		}

		if debug {
			fmt.Println("NEXT: ")
			fmt.Println()
		}
		for _, stateNeighbour := range getStateNeighbours(currentState) {
			if debug {
				PrintState(stateNeighbour.state)
			}
			newCost := costSoFar[currentState.String()] + stateNeighbour.cost

			if cost, ok := costSoFar[stateNeighbour.state.String()]; !ok || newCost < cost {
				costSoFar[stateNeighbour.state.String()] = newCost
				frontier.Enqueue(stateNeighbour.state, newCost)
				cameFrom[stateNeighbour.state.String()] = currentState.String()
			}
		}
	}

	return costSoFar[endState.String()]

}

func GetEndState(startState State) State {
	endState := make([]Pod, 0)
	endRow := startState.LowestHomeRow()
	startRow := hallRow + 1

	for podType, col := range homeColumns {
		for row := startRow; row <= endRow; row++ {
			pod := Pod{Point{row, col}, podType}
			endState = append(endState, pod)
		}

	}
	SortState(endState)
	return State(endState)
}

func getStateNeighbours(currentState State) []StateNeighbour {
	stateNeighbours := make([]StateNeighbour, 0)

	// next possible states are all the possible ways each amphipod can move
	for i, pod := range currentState {
		// get all possible next positions for this pod, and add those to state neighbours
		for nextPodPosition, cost := range GetPodNextPositionsAndCosts(pod, currentState) {
			nextState := make(State, len(currentState))
			copy(nextState, currentState)
			nextState[i] = nextPodPosition
			SortState(nextState[:])
			neighbour := StateNeighbour{nextState, cost}
			stateNeighbours = append(stateNeighbours, neighbour)
		}
	}

	return stateNeighbours
}

func GetPodNextPositionsAndCosts(pod Pod, state State) map[Pod]int {

	if !state.Contains(pod) {
		panic("current pod not in existing state")
	}

	nextPositions := make(map[Pod]int, 0)

	if pod.IsHome() && !pod.HomeButMustMakeSpace(state) {
		// home already, nowhere to go
		return nextPositions
	}

	if pod.InHallway() {
		// only place it can go is home
		homePos, homeCost, canReachHome := RouteHome(pod, state)
		if canReachHome {
			nextPositions[homePos] = homeCost
		}
		return nextPositions

	}

	// it's not home, it's not in hallway, it must be in starting position!

	// can't move if anyone above
	if pod.HasPodsAbove(state) {
		return nextPositions
	}

	// it can get into hallway

	// go as far left as possible
	for col := pod.Pt.Col; col >= hallStartCol; col-- {
		currentPosition := Point{hallRow, col}

		if col == homeColumns["A"] || col == homeColumns["B"] || col == homeColumns["C"] || col == homeColumns["D"] {
			// above home, cant stop here
			continue
		}
		if _, blocked := state.PodAt(currentPosition); blocked {
			break
		}

		nextPositions[Pod{currentPosition, pod.Type}] = DistanceCost(pod.Pt, currentPosition, pod.Type)

	}

	// go as far right as possible
	for col := pod.Pt.Col; col <= hallEndCol; col++ {
		currentPosition := Point{hallRow, col}

		if col == homeColumns["A"] || col == homeColumns["B"] || col == homeColumns["C"] || col == homeColumns["D"] {
			// above home, cant stop here
			continue
		}
		if _, blocked := state.PodAt(currentPosition); blocked {
			break
		}

		nextPositions[Pod{currentPosition, pod.Type}] = DistanceCost(pod.Pt, currentPosition, pod.Type)
	}

	// try go home
	homePos, homeCost, canReachHome := RouteHome(pod, state)
	if canReachHome {
		nextPositions[homePos] = homeCost
	}

	return nextPositions
}

func RouteHome(pod Pod, state State) (Pod, int, bool) {
	homePosition := state.GetHomePositions(pod.Type)

	// which direction should it go?
	goLeft := homePosition[0].Col < pod.Pt.Col

	var delta int
	if goLeft {
		delta = -1
	} else {
		delta = 1
	}

	// can it get home?
	targetCol := homePosition[0].Col
	for i := pod.Pt.Col; i != targetCol; i += delta {
		// move in direction of home
		currentPosition := Point{hallRow, i}
		if blockingPod, blocked := state.PodAt(currentPosition); blocked && blockingPod != pod {
			// blocked, can't get home
			return Pod{}, 0, false
		}
	}

	// it can get home! try all slots, starting from lowest
	for i := len(homePosition) - 1; i >= 0; i-- {
		homePos := homePosition[i]
		if _, hasPod := state.PodAt(homePos); !hasPod {
			return Pod{homePos, pod.Type}, DistanceCost(pod.Pt, homePos, pod.Type), true
		}
	}

	return Pod{}, 0, false
}

func (state State) GetHomePositions(podType string) []Point {
	startRow := hallRow + 1
	endRow := state.LowestHomeRow()
	col := homeColumns[podType]
	homePositions := make([]Point, 0)

	for row := startRow; row <= endRow; row++ {
		homePositions = append(homePositions, Point{row, col})
	}

	return homePositions

}

func (s State) PodAt(pt Point) (Pod, bool) {
	for _, pod := range s {
		if pod.Pt == pt {
			return pod, true
		}
	}
	return Pod{}, false
}

func (s State) ToMap() map[Point]string {
	stateMap := make(map[Point]string)
	for _, pod := range s {
		stateMap[pod.Pt] = pod.Type
	}
	return stateMap
}

func (pod Pod) HasPodsAbove(state State) bool {
	if pod.InHallway() {
		return false
	}

	if pod.Pt.Row == hallRow-1 {
		return false
	}

	positionAbovePod := Point{pod.Pt.Row - 1, pod.Pt.Col}
	if _, hasPodAbove := state.PodAt(positionAbovePod); hasPodAbove {
		return true
	}

	return false

}

func (pod Pod) HomeButMustMakeSpace(state State) bool {
	// if
	// - pod is home
	// - has no pods above it
	// - one pod below is wrong type
	// the pod is home, but it needs to let the other pod out
	if pod.IsHome() && !pod.HasPodsAbove(state) {
		col := pod.Pt.Col
		for row := pod.Pt.Row + 1; row <= state.LowestHomeRow(); row++ {
			position := Point{row, col}
			podAtPosition, hasPodAtPosition := state.PodAt(position)

			if !hasPodAtPosition {
				panic("HomeButMustMakeSpace: floating pod")
			}

			if podAtPosition.Type != pod.Type {
				// pod below of different time
				return true
			}
		}
	}
	return false
}

func (s State) LowestHomeRow() int {
	switch l := len(s); l {
	case smallState:
		return 3
	case largeState:
		return 5
	default:
		panic(fmt.Sprintf("unrecognized state size %v", l))
	}
}

func (p Pod) IsHome() bool {
	if !p.InHallway() && p.Pt.Col == homeColumns[p.Type] {
		// not in hallway, and in home column => home
		return true
	}
	return false
}

func (p Pod) InHallway() bool {
	return p.Pt.Row == hallRow
}

func DistanceCost(src, dst Point, podType string) int {
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

func (state State) String() string {
	var stateString strings.Builder
	stateMap := state.ToMap()
	for row := hallRow; row <= state.LowestHomeRow(); row++ {
		for col := hallStartCol; col <= hallEndCol; col++ {
			point := Point{row, col}
			podType, ok := stateMap[point]
			if ok {
				fmt.Fprint(&stateString, string(podType))
			} else if row == hallRow {
				fmt.Fprint(&stateString, ".")
			} else if col%2 != 0 && col != hallEndCol && col != hallStartCol {
				fmt.Fprint(&stateString, ".")
			} else {
				fmt.Fprint(&stateString, "#")
			}
		}
		fmt.Fprint(&stateString, "\n")
	}

	return stateString.String()

}

func PrintState(state State) {
	fmt.Println(state.String())
}

func PrintPath(startState State, cameFrom map[string]string) {
	current := GetEndState(startState).String()
	path := make([]string, 0)
	for current != startState.String() {
		path = append(path, current)
		current = cameFrom[current]
	}
	path = append(path, startState.String())

	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println()
		fmt.Println("STEP ", len(path)-i)
		fmt.Println()
		fmt.Println(path[i])
		fmt.Println()
	}

}
