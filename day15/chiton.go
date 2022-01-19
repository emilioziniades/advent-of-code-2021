package day15

import (
	"fmt"
	"math"
)

type point struct {
	x, y int
}

func Dijkstra(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	start := point{0, 0}
	end := point{h - 1, w - 1}
	nilNode := point{-1, -1}
	unvisited := make(map[point]bool)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			unvisited[point{x, y}] = true
		}
	}

	shortestPath := make(map[point]int)
	for node := range unvisited {
		shortestPath[node] = math.MaxInt
	}
	shortestPath[start] = 0
	prevNode := make(map[point]point)

	for len(unvisited) > 0 {
		currentMinNode := nilNode

		// finds node with lowest score
		for node := range unvisited {
			if currentMinNode == nilNode {
				currentMinNode = node
			} else if shortestPath[node] < shortestPath[currentMinNode] {
				currentMinNode = node
			}
		}

		neighbours := getNeighbours(currentMinNode, h, w)
		for _, n := range neighbours {
			tentativeValue := shortestPath[currentMinNode] + grid[n.y][n.x]
			if tentativeValue < shortestPath[n] {
				shortestPath[n] = tentativeValue
				prevNode[n] = currentMinNode
			}
		}
		delete(unvisited, currentMinNode)
	}
	printResult(prevNode, shortestPath, start, end)
	return shortestPath[end]
}

func getNeighbours(p point, h, w int) []point {
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
	reverse(path)
	fmt.Println(shortestPath[end], path)

}

func reverse[T any](t []T) {
	for i := len(t)/2 -1; i >= 0; i-- {
		opp := len(t) - 1 - i
		t[i], t[opp] = t[opp], t[i]
	}
}
