package day19

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

var rad = map[int]float64{
	0:   0,
	90:  (1.0 / 2.0) * math.Pi,
	180: math.Pi,
	270: (3.0 / 2.0) * math.Pi,
}

var rollTurnDict = make(map[int]string)

func init() {
	count := 0
	steps := ""
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ { // RTTT three times
			// R
			steps += "R"
			rollTurnDict[count] = steps
			count++
			// TTT
			for k := 0; k < 3; k++ {
				steps += "T"
				rollTurnDict[count] = steps
				count++
			}
		}

		// Do RTR
		steps += "RTR"
	}
}

// alpha: x-axis, beta: y-axis, gamma: z-axis
func rotate(p point, alpha, beta, gamma float64) point {
	pt := mat.NewDense(3, 1, []float64{
		float64(p.x),
		float64(p.y),
		float64(p.z),
	})

	sinA := math.Round(math.Sin(alpha))
	cosA := math.Round(math.Cos(alpha))
	sinB := math.Round(math.Sin(beta))
	cosB := math.Round(math.Cos(beta))
	sinC := math.Round(math.Sin(gamma))
	cosC := math.Round(math.Cos(gamma))

	Rx := mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, cosA, -1 * sinA,
		0, sinA, cosA,
	})

	Ry := mat.NewDense(3, 3, []float64{
		cosB, 0, sinB,
		0, 1, 0,
		-1 * sinB, 0, cosB,
	})

	Rz := mat.NewDense(3, 3, []float64{
		cosC, -1 * sinC, 0,
		sinC, cosC, 0,
		0, 0, 1,
	})

	// res = Rx * Ry * Rz * pt
	var res mat.Dense
	res.Mul(Rz, pt)
	res.Mul(Ry, &res)
	res.Mul(Rx, &res)
	data := res.RawMatrix().Data
	return point{int(data[0]), int(data[1]), int(data[2])}

}

func roll(p point) point {
	return rotate(p, rad[270], 0, 0)
}

func turn(p point) point {
	return rotate(p, 0, 0, rad[90])
}

func unroll(p point) point {
	return rotate(p, rad[90], 0, 0)
}

func unturn(p point) point {
	return rotate(p, 0, 0, rad[270])
}

// can arrive at all possible orientations via the sequence RTTTRTTTRTTT, then do (RTR), and then again RTTTRTTTRTTT
func possibleOrientations(pts []point) [][]point {
	res := make([][]point, 0)
	for i := 0; i < 24; i++ {
		res = append(res, make([]point, 0))
	}

	for _, p := range pts {
		c := 0
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ { // RTTT three times
				// R
				p = roll(p)
				res[c] = append(res[c], p)
				c++
				// TTT
				for k := 0; k < 3; k++ {
					p = turn(p)
					res[c] = append(res[c], p)
					c++
				}
			}

			// Do RTR
			p = roll(turn(roll(p)))
		}
	}

	return res
}

func unrollAndUnturn(p point, steps string) point {
	res := p
	n := len(steps) - 1
	for i := 0; i <= n; i++ {
		switch steps[i] {
		case 'R':
			res = unroll(res)
		case 'T':
			res = unturn(res)
		}
	}
	return res
}

func rollAndTurn(p point, steps string) point {
	res := p
	n := len(steps) - 1
	for i := 0; i <= n; i++ {
		switch steps[i] {
		case 'R':
			res = roll(res)
		case 'T':
			res = turn(res)
		}
	}
	return res
}

func format(m *mat.Dense) fmt.Formatter {
	return mat.Formatted(m, mat.Squeeze())

}
