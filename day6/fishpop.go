package day6

import (
	"github.com/emilioziniades/adventofcode2021/parse"
)

type popD map[int]int

func (pd popD) count() (i int) {
	for _, n := range pd {
		i += n
	}
	return
}

const (
	rate  = 6
	nrate = 8
)

func FishPopDict(initial []string, days int) int {
	p := parse.CommaSeparatedNumbers(initial)
	var pd popD = make(map[int]int)

	// initialize popD
	for _, fish := range p {
		pd[fish] += 1
	}

	// sim each day
	for i := 1; i <= days; i++ {
		births := pd[0]
		for i := 0; i <= 8; i++ {
			pd[i] = pd[i+1]
		}
		pd[6] += births
		pd[8] = births
	}
	return pd.count()
}

func FishPop(initial []string, days int) int {
	p := parse.CommaSeparatedNumbers(initial)

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
