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

func ParseState(file string) State {

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

	return state

}
