package day20

func Enhance(input []string, n int) int {
	pad := n + 1
	image, rule := parseImageAndRule(input, pad)
	current := image
	next := make([][]int, 0)
	for i := 0; i < n; i++ {
		next = enhance(current, rule)
		current = next

	}
	return countOnes(current)
}

func enhance(image [][]int, rule string) [][]int {
	newimage := makeGrid(len(image), len(image[0]))
	RC := []int{-1, 0, 1}

	// need to set all boundary points, but they need special treatment
	// can't do this normally because they don't have a 3 x 3 around them
	// but, their values will be uniform, all 0s or all 1s, and so their
	// value can be determined by the first and last indexes of the rule
	var zeroval int
	var oneval int
	if rule[0] == '#' {
		zeroval = 1
	}
	if rule[len(rule)-1] == '#' {
		oneval = 1
	}

	for r := range image {
		for c := range image[0] {
			if r == 0 || r == len(image)-1 || c == 0 || c == len(image[0])-1 {
				// pixel on boundary
				switch image[r][c] {
				case 1:
					newimage[r][c] = oneval
				case 0:
					newimage[r][c] = zeroval
				}

			} else {
				// pixel not on boundary, get 3 x 3 surrounding this point
				pixel := make([]int, 0)
				for _, rr := range RC {
					for _, cc := range RC {
						pixel = append(pixel, image[r+rr][c+cc])
					}
				}
				newimage[r][c] = pixelToInt(pixel, rule)
			}

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
	if rule[n] == '#' {
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
				panic("countOnes: pixel neither 1 or 0")
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
			switch e {
			case '#':
				grid[row] = append(grid[row], 1)
			case '.':
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
