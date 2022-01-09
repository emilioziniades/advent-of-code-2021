package day9

import (
	"sort"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

type point struct {
	x, y int
}

func LowPoints(fileSlice []string) (int, []point) {
	heatmap := parseInputToHeatmap(fileSlice)
	height, width := len(heatmap), len(heatmap[0])
	risk := 0
	lowpoints := make([]point, 0)

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
				lowpoints = append(lowpoints, point{x, y})
				risk += current + 1
			}
		}
	}
	return risk, lowpoints
}

func CountBasins(file []string) int {
	_, lowpoints := LowPoints(file)
	hm := parseInputToHeatmap(file)
	basins := make([]int, 0)

	for _, lowpoint := range lowpoints {
		basin := countBasin(hm, lowpoint.x, lowpoint.y)
		basins = append(basins, basin)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}
func countBasin(hm [][]int, x, y int) int {
	h, w := len(hm), len(hm[0])
	visited := make(map[point]bool)

	var countRec func(int, int)
	countRec = func(x, y int) {
		if visited[point{x, y}] || hm[y][x] == 9 {
			return
		}
		visited[point{x, y}] = true

		// has top
		if !(y == 0) {
			countRec(x, y-1)
		}
		// has bottom
		if !(y == h-1) {
			countRec(x, y+1)
		}
		// has left
		if !(x == 0) {
			countRec(x-1, y)
		}
		// has right
		if !(x == w-1) {
			countRec(x+1, y)
		}

		return
	}

	countRec(x, y)
	return len(visited)
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
