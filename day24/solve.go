package day24

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/emilioziniades/adventofcode2021/stack"
)

type Values struct {
	A, B, C int
}

type Single struct {
	Index  int
	Values Values
}

type Pair struct {
	Hit, Miss Single
}

func FindMaxValidModelNumber(name string) int {
	modelNumbers := FindAllValidModelNumbers(name)
	n := 0
	for _, e := range modelNumbers {
		if e > n {
			n = e
		}
	}
	return n
}

func FindMinValidModelNumber(name string) int {
	modelNumbers := FindAllValidModelNumbers(name)
	n := math.MaxInt
	for _, e := range modelNumbers {
		if e < n {
			n = e
		}
	}
	return n
}

func FindAllValidModelNumbers(name string) []int {
	changingValues := GetChangingValues(name)
	pairs := GetRestrictedPairs(changingValues)
	validModelNumbers := findAllValidModelNumbers(pairs)
	return validModelNumbers

}

func findAllValidModelNumbers(pairs []Pair) []int {
	modelNumberLength := len(pairs) * 2
	modelNumbers := make([][]int, 0)
	zeros := make([]int, modelNumberLength)
	modelNumbers = append(modelNumbers, zeros)

	for _, pair := range pairs {
		newModelNumbers := make([][]int, 0)

		for _, modelNumber := range modelNumbers {
			for missDigit := 1; missDigit <= 9; missDigit++ {
				hitDigit := missDigit + pair.Miss.Values.B + pair.Hit.Values.A

				if hitDigit >= 1 && hitDigit <= 9 {
					// this is a valid pair, insert values at those indexes
					newModelNumber := make([]int, modelNumberLength)
					copy(newModelNumber, modelNumber)
					newModelNumber[pair.Miss.Index] = missDigit
					newModelNumber[pair.Hit.Index] = hitDigit
					newModelNumbers = append(newModelNumbers, newModelNumber)
				}
			}
		}
		oldLen := len(modelNumbers)
		modelNumbers = append(modelNumbers, newModelNumbers...)
		modelNumbers = modelNumbers[oldLen:]
	}

	modelNumberInts := make([]int, 0)
	for _, s := range modelNumbers {
		modelNumberInts = append(modelNumberInts, IntSliceToInt(s))
	}

	return modelNumberInts
}

func IntSliceToInt(s []int) int {
	n := 0
	pow := 1

	for i := len(s) - 1; i >= 0; i-- {
		x := s[i]
		if x > 9 {
			panic("this only works for single digit slice values")
		}
		n += x * pow
		pow *= 10
	}

	return n
}

func GetRestrictedPairs(changingValues []Values) []Pair {
	restrictedPairs := make([]Pair, 0)
	missStack := stack.New[Single]()

	for i, vals := range changingValues {
		currentSingle := Single{Index: i, Values: vals}
		if vals.C == 1 {
			// this is a miss, add to stack
			missStack.Push(currentSingle)
		} else if vals.C == 26 {
			// this is a hit, get a miss off stack and put in pairs list
			missSingle := missStack.Pop()
			currentPair := Pair{Hit: currentSingle, Miss: missSingle}
			restrictedPairs = append(restrictedPairs, currentPair)
		}
	}

	return restrictedPairs

}

func GetChangingValues(filename string) []Values {
	changingValues := make([]Values, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	for {
		values := Values{}
		_, err := fmt.Fscanf(file, instructionStep, &values.C, &values.A, &values.B)
		if err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			log.Fatal(err)

		}
		changingValues = append(changingValues, values)
	}

	return changingValues
}
