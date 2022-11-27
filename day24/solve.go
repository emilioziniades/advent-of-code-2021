package day24

import (
	"fmt"
	"io"
	"log"
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

func FindAllModelNumbers(pairs []Pair) [][]int {
	modelNumberLength := len(pairs) * 2
	modelNumbers := make([][]int, 0)
	zeros := make([]int, modelNumberLength)
	modelNumbers = append(modelNumbers, zeros)

	for _, pair := range pairs {
		newModelNumbers := make([][]int, 0)
		copy(newModelNumbers, modelNumbers)
		for _, modelNumber := range modelNumbers {
			for i := 1; i <= 9; i++ {
				j := i + pair.Miss.Values.B + pair.Hit.Values.A
				newModelNumber := make([]int, modelNumberLength)
				copy(newModelNumber, modelNumber)
				if j >= 1 && j <= 9 {
					// this is a valid pair, insert it
					newModelNumber[pair.Miss.Index] = i
					newModelNumber[pair.Hit.Index] = j
					newModelNumbers = append(newModelNumbers, newModelNumber)
				}
			}
		}
		copy(modelNumbers, newModelNumbers)
	}

	return modelNumbers

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
