package day11

import "fmt"

func FlashCount(grid [][]int, steps int) int {
	fmt.Println(grid, len(grid), cap(grid))
	printGrid(grid)
	h, w := len(grid), len(grid[0])
	for i := 0; i < steps; i++ {
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				grid[r][c]++
			}
		}
		printGrid(grid)
	}
	return 0
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println("")
	}
}
