package day20

import (
	"fmt"
	"log"
)

func Enhance(input []string, n int) int {
	//	for j := 0; j < 40; j++ {
	image, rule := parseImageAndRule(input, 300)
	//		fmt.Println(len(image), len(image[0]))
	current := image
	//		fmt.Println(rule)
	//		fmt.Println("\nbefore enhancing: ")
	//		util.PrintGrid(current)
	var next [][]int
	for i := 0; i < 2; i++ {
		next = enhance(current, rule)
		//			fmt.Printf("enhancing, round %d\n", i+1)
		//			util.PrintGrid(next)
		//		return 35
		current = next

	}
	fmt.Printf("with %d padding, result is %d\n", 300, countOnes(current))
	//	}
	return countOnes(current)
}

func enhance(image [][]int, rule string) [][]int {
	newimage := makeGrid(len(image), len(image[0]))
	rstart := 1
	cstart := 1
	rend := len(image) - 2
	cend := len(image[0]) - 2

	RC := []int{-1, 0, 1}

	for r := rstart; r <= rend; r++ {
		for c := cstart; c <= cend; c++ {
			// get 3 x 3 surrounding this point
			pixel := make([]int, 0)
			for _, rr := range RC {
				for _, cc := range RC {
					pixel = append(pixel, image[r+rr][c+cc])
				}
			}
			newimage[r][c] = pixelToInt(pixel, rule)

		}
	}
	return newimage
}

func pixelToInt(pixel []int, rule string) int {
	n := 0
	l := len(pixel) - 1
	for i, e := range pixel {
		bit := l - i
		n += e << bit
	}
	if s := string(rule[n]); s == "#" {
		return 1
	} else {
		return 0
	}
}

func countOnes(image [][]int) (count int) {
	for r := 0; r < len(image); r++ {
		for c := 0; c < len(image[0]); c++ {
			switch image[r][c] {
			case 1:
				count++
			case 0:
				continue
			default:
				log.Fatal("countOnes: something has gone wrong")
			}
		}
	}
	return
}

func parseImageAndRule(input []string, padding int) ([][]int, string) {
	rule := input[0]
	grid := make([][]int, len(input)-2)
	row := 0
	for i, line := range input {
		if i == 0 || line == "" {
			continue
		}
		grid[row] = make([]int, 0)
		for _, e := range line {
			switch string(e) {
			case "#":
				grid[row] = append(grid[row], 1)
			case ".":
				grid[row] = append(grid[row], 0)
			}
		}
		row++
	}

	//padding
	rowpad := len(grid[0]) + 2*padding

	// pad columns
	for i := 0; i < len(grid); i++ {
		grid[i] = append(make([]int, padding), grid[i]...)
		grid[i] = append(grid[i], make([]int, padding)...)
	}

	// pad rows
	for i := 0; i < padding; i++ {
		grid = append([][]int{make([]int, rowpad)}, grid...)
		grid = append(grid, make([]int, rowpad))
	}

	return grid, rule

}

func makeGrid(rows, cols int) [][]int {
	g := make([][]int, rows)
	for i := range g {
		g[i] = make([]int, cols)
	}
	return g
}
