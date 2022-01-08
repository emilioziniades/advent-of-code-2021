package day9

import (
	"sort"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

func LowPoints(fileSlice []string) int {
	heatmap := parseInputToHeatmap(fileSlice)
	height, width := len(heatmap), len(heatmap[0])
	risk := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			current := heatmap[y][x]
			neighbours := make([]int, 0)

			//not top side, so has top neighbour
			if !(y == 0) {
				neighbours = append(neighbours, heatmap[y-1][x])
			}

			//not bottom side, so has bottom neighbour
			if !(y == height-1) {
				neighbours = append(neighbours, heatmap[y+1][x])
			}

			//not left side, so has left neighbour
			if !(x == 0) {
				neighbours = append(neighbours, heatmap[y][x-1])
			}

			//not right side, so has right neighbour
			if !(x == width-1) {
				neighbours = append(neighbours, heatmap[y][x+1])
			}

			sort.Ints(neighbours)
			if current < neighbours[0] {
				risk += current + 1
			}
		}
	}
	return risk
}

func parseInputToHeatmap(fs []string) [][]int {
	heatmap := make([][]int, len(fs))
	for i, e := range fs {
		digitsString := strings.Split(e, "")
		digits, _ := parse.StringToIntSlice(digitsString)
		heatmap[i] = append(heatmap[i], digits...)
	}
	return heatmap
}
