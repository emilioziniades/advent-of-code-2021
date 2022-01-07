package day7

import (
	"math"
	"sort"
)

// MinCost finds the value of x in range min(nums), max(nums)
// which minimizes cost of shifting each int in nums
// where cost_n is determined by costfunc, the cost function
// and cost = sum(0, n) { cost_n }
func MinCost(nums []int, costfunc func(float64) float64) int {
	t := math.Inf(1)
	min, max := minMax(nums)
	for i := min; i <= max; i++ {
		if c := cost(nums, i, costfunc); c < t {
			t = c
		}
	}
	return int(t)

}

func cost(nums []int, target int, costfunc func(float64) float64) float64 {
	c := float64(0)
	for _, num := range nums {
		dist := math.Abs(float64(num - target))
		c += costfunc(dist)
	}
	return c
}

func CostConst(n float64) float64 { return n }

func CostTriangle(n float64) float64 { return (n * (n + 1)) / 2 }

func minMax(x []int) (int, int) {
	sort.Ints(x)
	return x[0], x[len(x)-1]
}
