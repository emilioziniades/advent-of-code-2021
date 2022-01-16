package day12

import (
	"fmt"
	"strings"
)

func CountPaths(in []string) int {
	paths := pathsFromStringSlice(in)
	fmt.Println(paths)
	return 0
}

func pathsFromStringSlice(in []string) map[string][]string {
	gamma := make(map[string][]string)
	for _, e := range in {
		nodes := strings.Split(e, "-")
		// assuming lines of the form A-B (only two nodes joined by '-')
		gamma[nodes[0]] = append(gamma[nodes[0]], nodes[1])
		gamma[nodes[1]] = append(gamma[nodes[1]], nodes[0])
	}
	return gamma
}

func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}
