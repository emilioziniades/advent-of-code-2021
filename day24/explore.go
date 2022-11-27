package day24

import (
	"fmt"
	"strings"
)

var (
	instructionStep = `inp w
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
)

func Step(z, i, a, b, c int) int {
	if i-a == z%26 {
		return z / c
	} else {
		return 26*(z/c) + i + b
	}
}

func StepAll(input int, changingValues []Values) int {
	digits, skip := GetDigits(input)
	if skip {
		return -1
	}

	var z int

	for idx, inp := range digits {
		vals := changingValues[idx]
		z = Step(z, inp, vals.A, vals.B, vals.C)
	}

	return z
}

func RemakeInstructions(changingValues []Values) string {
	instructions := strings.Builder{}

	for _, vals := range changingValues {
		step := fmt.Sprintf(instructionStep, vals.C, vals.A, vals.B)
		instructions.WriteString(step)
	}

	return instructions.String()
}
