package day7

import (
	"math"
)

// MinCost finds the value of x in range min(i), max(i)
// which minimizes cost of shifting each int in i
// where cost_n = | i_n - x |
// and cost = sum(0, n) { cost_n }
func MinCost(nums []int) int {
	t := math.Inf(1)
	min, max := findMinMax(nums)
	for i := min; i <= max; i++ {
		if c := cost(nums, i); c < t {
			t = c
		}
	}
	return int(t)

}

func cost(nums []int, target int) float64 {
	c := float64(0)
	for _, num := range nums {
		c += math.Abs(float64(num - target))
	}
	return c
}

func findMinMax(nums []int) (int, int) {
	min := math.Inf(1)
	max := math.Inf(-1)
	for _, num := range nums {
		n := float64(num)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

	}
	return int(min), int(max)
}
