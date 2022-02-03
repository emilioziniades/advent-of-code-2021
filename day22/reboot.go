package day22

import (
	"math"
	"regexp"
	"strconv"
)

type point struct {
	x, y, z int
}

// type cuboid struct {
// SSS, SSE, SES, SEE, ESS, ESE, EES, EEE point
// }

// func (c cuboid) Volume() int {
// return int(math.Pow(c.SSS-c.SSE+1, 3))
// }

func Reboot(input []string, bounds float64) int {
	reactor := make(map[point]bool)
	re := regexp.MustCompile(`([a-z]+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+).*?(-?\d+)`)
	for _, line := range input {
		if re.MatchString(line) {
			match := re.FindStringSubmatch(line)
			on := false
			if match[1] == "on" {
				on = true
			}
			xS := intInRange(match[2], bounds, true)
			xE := intInRange(match[3], bounds, false)
			yS := intInRange(match[4], bounds, true)
			yE := intInRange(match[5], bounds, false)
			zS := intInRange(match[6], bounds, true)
			zE := intInRange(match[7], bounds, false)

			// currCuboid := cuboid{
			// point{xS, yS, zS},
			// point{xS, yS, zE},
			// point{xS, yE, zS},
			// point{xS, yE, zE},
			// point{xE, yS, zS},
			// point{xE, yS, zE},
			// point{xE, yE, zS},
			// point{xE, yE, zE},
			// }

			for x := xS; x <= xE; x++ {
				for y := yS; y <= yE; y++ {
					for z := zS; z <= zE; z++ {
						// fmt.Printf("%d,%d,%d\n", x, y, z)
						reactor[point{x, y, z}] = on

					}
				}
			}
		} else {
			panic("parsing error")
		}
	}
	return countOn(reactor)
}

func countOn(r map[point]bool) (on int) {
	for _, o := range r {
		if o == true {
			on++
		}
	}
	return
}

func intInRange(s string, r float64, start bool) int {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	if start {
		return int(math.Max(f, -1*r))
	} else {
		return int(math.Min(f, r))
	}
}
