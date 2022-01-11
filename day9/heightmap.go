package day9

import (
	"sort"
)

type point struct {
	x, y int
}

func LowPoints(grid [][]int) (int, []point) {
	height, width := len(grid), len(grid[0])
	risk := 0
	lowpoints := make([]point, 0)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			current := grid[y][x]
			neighbours := make([]int, 0)

			//not top side, so has top neighbour
			if !(y == 0) {
				neighbours = append(neighbours, grid[y-1][x])
			}

			//not bottom side, so has bottom neighbour
			if !(y == height-1) {
				neighbours = append(neighbours, grid[y+1][x])
			}

			//not left side, so has left neighbour
			if !(x == 0) {
				neighbours = append(neighbours, grid[y][x-1])
			}

			//not right side, so has right neighbour
			if !(x == width-1) {
				neighbours = append(neighbours, grid[y][x+1])
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

func CountBasins(grid [][]int) int {
	_, lowpoints := LowPoints(grid)
	basins := make([]int, 0)

	for _, lowpoint := range lowpoints {
		basin := countBasin(grid, lowpoint)
		basins = append(basins, basin)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}
func countBasin(grid [][]int, p point) int {
	h, w := len(grid), len(grid[0])
	visited := make(map[point]bool)

	var countRec func(point)
	countRec = func(p point) {
		if visited[p] || grid[p.y][p.x] == 9 {
			return
		}
		visited[p] = true

		// has top
		if !(p.y == 0) {
			countRec(point{p.x, p.y - 1})
		}
		// has bottom
		if !(p.y == h-1) {
			countRec(point{p.x, p.y + 1})
		}
		// has left
		if !(p.x == 0) {
			countRec(point{p.x - 1, p.y})
		}
		// has right
		if !(p.x == w-1) {
			countRec(point{p.x + 1, p.y})
		}

		return
	}

	countRec(p)
	return len(visited)
}
