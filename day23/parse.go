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

	State [8]Pod
)

// implement sort.Interface for State
func (s State) Len() int {
	return len(s)
}

func (s State) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s State) Less(i, j int) bool {
	return s[i].Id() < s[j].Id()
}

// unique ID for that point
func (p Pod) Id() int {
	return p.Pt.Row*maxRows + p.Pt.Col
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

	state := State{}

	data, err := parse.FileToStringSlice(file)

	if err != nil {
		log.Fatalln(err)
	}

	index := 0

	for r, row := range data {
		for c, char := range row {
			if letter := string(char); strings.Contains("ABCD", letter) {
				current := Pod{Point{r, c}, letter}
				state[index] = current
				index++
			}
		}
	}

	sort.Sort(state)

	return state

}
