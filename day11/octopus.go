package day11

import (
	"fmt"

	"github.com/emilioziniades/adventofcode2021/queue"
)

var R = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var C = []int{-1, 0, 1, -1, 1, -1, 0, 1}

type point struct {
	R, C int
}

func FlashCount(grid [][]int, steps int) (int, int) {
	count := 0
	syncStep := 10000000000
	fmt.Println("Before any steps:")
	printGrid(grid)
	h, w := len(grid), len(grid[0])
	for i := 1; i <= steps; i++ {
		// Step 1 : increase every octopus' energy level by 1
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				grid[r][c]++
			}
		}

		// Step 2 : Any octopus > 9 flashes. Neighbours increase by 1. Also flash if > 9
		flashed := make(map[point]bool)
		q := queue.New[point]()
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if grid[r][c] > 9 {
					p := point{r, c}
					q.Enqueue(p)
					flashed[p] = true
				}
			}
		}

		for len(q) > 0 {
			curr := q.Dequeue()
			// flash neighbours
			for n := 0; n < 8; n++ {
				row, col := curr.R+R[n], curr.C+C[n]
				if row >= 0 && row < h && col >= 0 && col < w {
					grid[row][col]++
					if grid[row][col] > 9 && !flashed[point{row, col}] {
						p := point{row, col}
						q.Enqueue(p)
						flashed[p] = true
					}
				}
			}
		}
		// Step 3: Any octopus that flashed is set to 0
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				if grid[r][c] > 9 {
					count++
					grid[r][c] = 0

				}
			}
		}

		if i%5 == 0 {
			fmt.Printf("After step %d:\n", i)
			printGrid(grid)
		}

		// checks if all octopi flashed in a single step.
		// syncStep should only record the first instance of this
		if len(flashed) == w*h && i < syncStep {
			syncStep = i
		}
	}
	return count, syncStep
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
