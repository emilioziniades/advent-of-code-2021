package day24

import (
	"fmt"
	"strings"
)

var (
	changingValues = [14][3]int{
		{10, 2, 1},
		{14, 13, 1},
		{14, 13, 1},
		{-13, 9, 26},
		{10, 15, 1},
		{-13, 3, 26},
		{-7, 6, 26},
		{11, 5, 1},
		{10, 16, 1},
		{13, 1, 1},
		{-4, 6, 26},
		{-9, 3, 26},
		{-13, 7, 26},
		{-9, 9, 26},
	}
)

func RemakeInstructions() string {
	instructions := strings.Builder{}
	instructionStep := `inp w
mul x 0
add x z
mod x 26
div z %v
add x %v
eql x w
eql x 0
mul y 0
add y 25
mul y x
add y 1
mul z y
mul y 0
add y w
add y %v
mul y x
add z y
`

	for _, vals := range changingValues {
		a, b, dz := vals[0], vals[1], vals[2]
		step := fmt.Sprintf(instructionStep, dz, a, b)
		instructions.WriteString(step)
	}

	return instructions.String()
}

func Step(z, i, a, b, dz int) int {
	if i-a == z%26 {
		return z / dz
	} else {
		return 26*(z/dz) + i + b
	}
}

func StepAll(input int) int {
	digits, skip := GetDigits(input)
	if skip {
		return -1
	}

	var z int

	for idx, inp := range digits {
		vals := changingValues[idx]
		a, b, dz := vals[0], vals[1], vals[2]
		z = Step(z, inp, a, b, dz)
	}

	return z
}
