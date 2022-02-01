package day15

import (
	"fmt"
	"math"

	"github.com/emilioziniades/adventofcode2021/queue"
	"github.com/emilioziniades/adventofcode2021/util"
)

type point struct {
	x, y int
}

// implementing util.Vector interface
func (p point) ToSlice() []int {
	return []int{p.x, p.y}
}

func ShortestPath(grid [][]int, heuristicFunc func(util.Vector, util.Vector) int) int {
	h, w := len(grid), len(grid[0])
	start := point{0, 0}
	end := point{h - 1, w - 1}
	prevNode := make(map[point]point)
	shortestPath := make(map[point]int)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			shortestPath[point{x, y}] = math.MaxInt
		}
	}
	shortestPath[start] = 0

	q := queue.NewPriority[point]()
	q.Enqueue(start, 0)

	for q.Len() > 0 {
		currentMinNode := q.Dequeue().Value

		if currentMinNode == end {
			break
		}

		for _, n := range neighbours(currentMinNode, h, w) {
			cost := grid[n.y][n.x]
			tentativeValue := shortestPath[currentMinNode] + cost
			if tentativeValue < shortestPath[n] {
				shortestPath[n] = tentativeValue
				prevNode[n] = currentMinNode
				priority := tentativeValue + heuristicFunc(n, end)
				q.Enqueue(n, priority)
			}
		}
	}
	return shortestPath[end]
}

func Dijkstra(grid [][]int) int {
	noHeuristic := func(a, b util.Vector) int { return 0 }
	return ShortestPath(grid, noHeuristic)
}

func AStar(grid [][]int) int {
	return ShortestPath(grid, util.ManhattanDistance)
}

func AStarFivefold(grid [][]int) int {
	bigGrid := extendGrid(grid)
	return AStar(bigGrid)
}

func neighbours(p point, h, w int) []point {
	Y := []int{0, 1, 0, -1}
	X := []int{1, 0, -1, 0}
	neighbours := make([]point, 0)

	for i := 0; i < 4; i++ {
		x, y := p.x+X[i], p.y+Y[i]
		if y >= 0 && y < h && x >= 0 && x < w {
			neighbours = append(neighbours, point{x, y})
		}
	}
	return neighbours
}

func printResult(prevNode map[point]point, shortestPath map[point]int, start, end point) {
	path := make([]point, 0)
	node := end

	for node != start {
		path = append(path, node)
		node = prevNode[node]
	}
	path = append(path, start)
	util.Reverse(path)
	fmt.Println(shortestPath[end], path)

}

func extendGrid(grid [][]int) [][]int {
	h, _ := len(grid), len(grid[0])
	res := make([][]int, h)

	// fill horizontally first
	for i := 0; i < h; i++ {
		for j := 0; j < 5; j++ {

			curr := util.Map(grid[i], func(n int) int { return (n+j-1)%9 + 1 })
			res[i] = append(res[i], curr...)
		}
	}

	// then fill vertically
	for i := 1; i < 5; i++ {
		for j := 0; j < h; j++ {
			curr := util.Map(res[j], func(n int) int { return (n+i-1)%9 + 1 })
			res = append(res, curr)
		}
	}
	return res
}
