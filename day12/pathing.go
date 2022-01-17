package day12

import (
	"fmt"
	"strings"
)

func CountPaths(in []string, hasFunc func([]string, string) bool) int {
	graph := graphFromStringSlice(in)
	paths := make([][]string, 0)

	var recPaths func([]string)
	recPaths = func(s []string) {
		curr := s[len(s)-1]
		if curr == "end" {
			paths = append(paths, s)
			return
		}

		for _, e := range graph[curr] {
			if e == "start" {
				continue
			}
			if !hasFunc(s, e) || IsUpper(e) {
				recPaths(append(s, e))
			}
		}
	}

	recPaths([]string{"start"})
	return len(paths)
}

func CountPathsOne(in []string) int {
	return CountPaths(in, has)
}

func CountPathsTwo(in []string) int {
	return CountPaths(in, hasAndLowerTwice)
}

func has(sl []string, s string) bool {
	for _, e := range sl {
		if e == s {
			return true
		}
	}
	return false
}

// hasAndLowerTwice should return true if:
// - any lowercase cave appears twice AND s already included in sl
// otherwise return false
func hasAndLowerTwice(sl []string, s string) bool {

	c := make(map[string]int)
	var twice, has bool

	for _, e := range sl {
		if IsLower(e) {
			c[e]++
		}
		if e == s {
			has = true
		}
	}

	// checks if any lowercase cave visited twice
	for _, v := range c {
		if v >= 2 {
			twice = true
		}
	}

	return twice && has
}

func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}

func graphFromStringSlice(in []string) map[string][]string {
	gamma := make(map[string][]string)
	for _, e := range in {
		nodes := strings.Split(e, "-")
		// assuming lines of the form A-B (only two nodes joined by '-')
		gamma[nodes[0]] = append(gamma[nodes[0]], nodes[1])
		gamma[nodes[1]] = append(gamma[nodes[1]], nodes[0])
	}
	return gamma
}

func printPath(s []string) {
	fmt.Println(strings.Join(s, ","))
}
