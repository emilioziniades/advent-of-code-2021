package day8

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
}

func UniqueDigits(entries []string) (n int) {
	for _, entry := range entries {
		_, output := parseEntry(entry)
		c := countUnique(output)
		n += c
	}
	return
}

func parseEntry(s string) ([]string, []string) {
	spl := strings.Split(s, " ")
	pattern, output := make([]string, 0), make([]string, 0)
	pipe := false

	for _, sp := range spl {
		if sp == "|" {
			pipe = true
			continue
		}
		if !pipe {
			pattern = append(pattern, sp)
		} else {
			output = append(output, sp)
		}
	}
	//	fmt.Println(pattern, output)
	return pattern, output
}

func countUnique(pattern []string) int {
	n := 0
	for _, p := range pattern {
		if l := len(p); ((l >= 2) && (l <= 4)) || l == 7 {
			//			fmt.Println(p)
			n += 1
		}
	}
	return n
}
