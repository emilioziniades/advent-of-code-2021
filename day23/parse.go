// parses input into graph
package day23

import (
	"log"
	"sort"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

type (
	Pod struct {
		Pt   Point
		Type string
	}

	Point struct {
		Row, Col int
	}

	State []Pod
)

func SortState(s []Pod) {
	sort.Slice(s, func(i, j int) bool {
		if s[i].Pt.Row != s[j].Pt.Row {
			return s[i].Pt.Row < s[j].Pt.Row
		} else {
			return s[i].Pt.Col < s[j].Pt.Col
		}
	})
}

func (s State) Contains(p Pod) bool {
	for _, e := range s {
		if p == e {
			return true
		}
	}
	return false
}

func ParseState(file string, withExtraState bool) State {

	state := make(State, 0)

	data, err := parse.FileToStringSlice(file)

	if err != nil {
		log.Fatalln(err)
	}

	for r, row := range data {
		for c, char := range row {
			if letter := string(char); strings.Contains("ABCD", letter) {
				current := Pod{Point{r, c}, letter}
				state = append(state, current)
			}
		}
	}

	SortState(state)

	if withExtraState {
		return AddMoreState(state)
	}

	return state
}

func AddMoreState(state State) State {
	newState := make(State, len(state))
	copy(newState, state)

	// pods in row 3 go to row 5
	for i, pod := range newState {
		if pod.Pt.Row == 3 {
			pod.Pt.Row = 5
			newState[i] = pod
		}
	}

	// insert the following pods into state
	podsToAdd := []Pod{
		{Point{3, 3}, "D"},
		{Point{4, 3}, "D"},
		{Point{3, 5}, "C"},
		{Point{4, 5}, "B"},
		{Point{3, 7}, "B"},
		{Point{4, 7}, "A"},
		{Point{3, 9}, "A"},
		{Point{4, 9}, "C"},
	}

	newState = append(newState, podsToAdd...)

	SortState(newState)

	return newState

}
