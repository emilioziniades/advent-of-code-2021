package day14

import (
	"strings"
)

type pairCount struct {
	m          map[string]int
	begin, end string
}

func Polymer(in []string, steps int) int {
	t, rules := parseTemplateAndRules(in)
	pc := pairCount{m: make(map[string]int), begin: string(t[0]), end: string(t[len(t)-1])}

	// parse initial template into pair count
	for i := 0; i < len(t)-1; i++ {
		tt := t[i : i+2]
		pc.m[tt]++
	}

	// perform steps of polymerization by updating pc
	for i := 1; i <= steps; i++ {
		//copy pc.m to curr
		curr := make(map[string]int)
		for k, v := range pc.m {
			curr[k] = v
		}

		for k, v := range pc.m {
			childPairs := rules[k]
			cp1, cp2 := childPairs[0], childPairs[1]
			curr[cp1] += v
			curr[cp2] += v
			curr[k] -= v

		}
		pc.m = curr
	}

	//count letters from pc into letterCount (which includes duplicates)
	letterCount := make(map[string]int)
	for k, v := range pc.m {
		first, second := string(k[0]), string(k[1])
		letterCount[first] += v
		letterCount[second] += v
	}

	// deduplicate letterCount
	for k, v := range letterCount {
		if k == pc.begin || k == pc.end {
			letterCount[k] = (v-1)/2 + 1
		} else {
			letterCount[k] = v / 2
		}
	}
	return diffMinMax(letterCount)
}

func diffMinMax(m map[string]int) int {
	min, max := (1<<63)-1, 0
	for _, v := range m {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func parseTemplateAndRules(in []string) (string, map[string][]string) {
	templ := in[0]
	rules := make(map[string][]string)

	for i, e := range in {
		if i < 2 {
			continue
		}
		r := strings.Split(e, " -> ")
		rules[r[0]] = []string{string(r[0][0]) + r[1], r[1] + string(r[0][1])}
	}
	return templ, rules
}
