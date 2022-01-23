package day17

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func FindMaxY(t string) (int, int) {
	s, e := parseRange(t)
	lim := 1 << 11
	maxY := math.MinInt
	hitCount := 0
	for x := 1; x < lim; x++ {
		for y := -1 * lim; y < lim; y++ {
			v := point{x, y}
			y, hit := findMaxY(s, e, v)

			if hit {
				hitCount += 1
			}

			if y > maxY {
				maxY = y
			}
		}
	}
	return maxY, hitCount
}

func findMaxY(start, end, velocity point) (int, bool) {
	location := point{0, 0}
	maxY := math.MinInt
	for {
		if location.y > maxY {
			maxY = location.y
		}
		location.x += velocity.x
		location.y += velocity.y

		if velocity.x > 0 {
			velocity.x -= 1
		} else if velocity.x < 0 {
			velocity.x += 1
		}
		velocity.y -= 1

		if location.x >= start.x && location.x <= end.x && location.y <= start.y && location.y >= end.y {
			break
		}

		if location.x > end.x || location.y < end.y {
			return math.MinInt, false
		}
	}
	return maxY, true
}

func parseRange(t string) (point, point) {
	fmt.Println(t)
	t = strings.TrimPrefix(t, "target area: ")
	t = strings.Replace(t, ",", "", -1)
	t = strings.Replace(t, "x=", "", -1)
	t = strings.Replace(t, "y=", "", -1)
	t = strings.Replace(t, " ", "..", -1)
	ts := strings.Split(t, "..")
	tn := make([]int, 0)
	for _, e := range ts {
		n, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		tn = append(tn, n)
	}
	rStart := point{tn[0], tn[3]}
	rEnd := point{tn[1], tn[2]}
	fmt.Printf("start %v end %v\n", rStart, rEnd)
	return rStart, rEnd
}
