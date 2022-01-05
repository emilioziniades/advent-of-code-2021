package day6

import (
	"log"
	"strings"

	"github.com/emilioziniades/adventofcode2021/fetch"
)

type pop []int

const (
	rate  = 6
	nrate = 8
)

func (p pop) String() string {
	return strings.Join(fetch.IntToStringSlice(p), ",")
}

func SimDays(initial []string, days int) int {
	p := parseInitialPop(initial)

	// each day
	for i := 1; i <= days; i++ {
		// consider each fish
		for i := range p {
			if p[i] == 0 {
				p[i] = rate
				p = append(p, nrate)
			} else {
				p[i]--
			}
		}
	}
	return len(p)
}

func parseInitialPop(initial []string) pop {
	popI, err := fetch.StringToIntSlice(strings.Split(initial[0], ","))
	if err != nil {
		log.Fatalf("Converting inital string to pop: %s\n", err)
	}
	return popI
}
