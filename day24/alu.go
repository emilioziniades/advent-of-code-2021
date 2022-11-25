package day24

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/emilioziniades/adventofcode2021/parse"
)

const (
	debug    = false
	MinInput = 11111111111111
	MaxInput = 99999999999999
)

// const debug = true

func Run(program []string, input int, outVar string) int {
	inputs, skip := GetDigits(input)
	if skip {
		return -1
	}
	vars := map[string]int{
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	for _, step := range program {
		if debug {
			fmt.Println(vars)
			fmt.Println(step)
		}
		step := strings.Split(step, " ")
		op := step[0]
		switch op {
		case "inp":
			i := inputs[0]
			inputs = inputs[1:]
			s1 := step[1]
			vars[s1] = i

		case "add":
			s1 := step[1]
			s2 := step[2]
			if v2, ok := vars[s2]; ok {
				vars[s1] += v2
			} else {
				v2, _ := strconv.Atoi(s2)
				vars[s1] += v2
			}

		case "mul":
			s1 := step[1]
			s2 := step[2]
			if v2, ok := vars[s2]; ok {
				vars[s1] *= v2
			} else {
				v2, _ := strconv.Atoi(s2)
				vars[s1] *= v2
			}

		case "div":
			s1 := step[1]
			s2 := step[2]
			if v2, ok := vars[s2]; ok {
				if v2 == 0 {
					panic("div 0")
				}
				vars[s1] /= v2
			} else {
				v2, _ := strconv.Atoi(s2)
				vars[s1] /= v2
			}

		case "mod":
			s1 := step[1]
			s2 := step[2]

			if v2, ok := vars[s2]; ok {
				vars[s1] %= v2
			} else {
				v2, _ := strconv.Atoi(s2)
				vars[s1] %= v2
			}

		case "eql":
			v1 := step[1]
			v2 := step[2]
			if _, ok := vars[v2]; ok {
				if vars[v1] == vars[v2] {
					vars[v1] = 1
				} else {
					vars[v1] = 0
				}
			} else {
				v2, _ := strconv.Atoi(v2)
				if vars[v1] == v2 {
					vars[v1] = 1
				} else {
					vars[v1] = 0
				}
			}

		default:
			msg := fmt.Sprintln("unrecognized operation: ", op)
			panic(msg)

		}

	}
	return vars[outVar]
}

func LoadProgram(filename string) []string {
	rawProgram, err := parse.FileToStringSlice(filename)

	if err != nil {
		log.Fatalf("LoadProgram: %v", err)
	}

	// TODO : parse string into more useful structure

	return rawProgram

}

func ValidateModelNumber(filename string) int {
	program := LoadProgram(filename)
	for i := 99999999999999; i >= 11111111111111; i-- {
		// for i := 19999999999999; i <= 99999999999999; i += 10000000000000 {
		// for i := 11111111111111; i <= 99999999999999; i++ {
		n := Run(program, i, "z")
		fmt.Println(i, "\t", n)
		if n == 0 {
			return i
		}
		// break
	}
	return -1
}

func GetDigits(n int) ([]int, bool) {
	s := strconv.Itoa(n)
	strSlice := strings.Split(s, "")
	intSlice := make([]int, 0)

	if strings.Contains(s, "0") {
		return intSlice, true
	}

	for _, e := range strSlice {
		i, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, i)
	}

	return intSlice, false
}