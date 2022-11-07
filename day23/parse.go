// parses input into graph
package day23

import (
	"log"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

type (
	Point struct {
		Row, Col int
	}

	State map[Point]rune
)

func ParseInitialState(file string) State {

	state := make(State)

	data, err := parse.FileToStringSlice(file)

	if err != nil {
		log.Fatalln(err)
	}

	for r, row := range data {
		for c, char := range row {
			current := Point{r, c}
			if letter := string(char); strings.Contains("ABCD", letter) {
				state[current] = char
			}
		}
	}

	return state

}
