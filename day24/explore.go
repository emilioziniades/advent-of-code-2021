package day24

import (
	"fmt"
	"strings"
)

var ChangingValues = [14][3]int{
	{10, 2, 1},   // 0
	{14, 13, 1},  // 1
	{14, 13, 1},  // 2
	{-13, 9, 26}, // 3
	{10, 15, 1},  // 4
	{-13, 3, 26}, // 5
	{-7, 6, 26},  // 6
	{11, 5, 1},   // 7
	{10, 16, 1},  // 8
	{13, 1, 1},   // 9
	{-4, 6, 26},  // 10
	{-9, 3, 26},  // 11
	{-13, 7, 26}, // 12
	{-9, 9, 26},  // 13
}

func Step(z, i, a, b, c int) int {
	if i-a == z%26 {
		return z / c
	} else {
		return 26*(z/c) + i + b
	}
}

func StepAll(input int) int {
	digits, skip := GetDigits(input)
	if skip {
		return -1
	}

	var z int

	for idx, inp := range digits {
		vals := ChangingValues[idx]
		a, b, c := vals[0], vals[1], vals[2]
		z = Step(z, inp, a, b, c)
	}

	return z
}

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

	for _, vals := range ChangingValues {
		a, b, c := vals[0], vals[1], vals[2]
		step := fmt.Sprintf(instructionStep, c, a, b)
		instructions.WriteString(step)
	}

	return instructions.String()
}
