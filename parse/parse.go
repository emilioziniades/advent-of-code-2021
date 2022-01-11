package parse

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func FileToStringSlice(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]string, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		res = append(res, input.Text())
	}
	return res, nil
}

func FileToIntSlice(file string) ([]int, error) {
	stringRes, err := FileToStringSlice(file)
	if err != nil {
		return nil, err
	}
	intRes, err := StringToIntSlice(stringRes)
	if err != nil {
		return nil, err
	}
	return intRes, nil
}

func StringToIntSlice(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice))
	for i := range stringSlice {
		n, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			return nil, err
		}
		intSlice[i] = n
	}
	return intSlice, nil
}

func IntToStringSlice(intSlice []int) []string {
	strSlice := make([]string, len(intSlice))
	for i := range intSlice {
		s := strconv.Itoa(intSlice[i])
		strSlice[i] = s
	}
	return strSlice
}

func CommaSeparatedNumbers(s []string) []int {
	intSlice, err := StringToIntSlice(strings.Split(s[0], ","))
	if err != nil {
		log.Fatalf("Converting inital string to pop: %s\n", err)
	}
	return intSlice
}

func stringSliceToGrid(fs []string) [][]int {
	grid := make([][]int, len(fs))
	for i, e := range fs {
		digitsString := strings.Split(e, "")
		digits, _ := StringToIntSlice(digitsString)
		grid[i] = append(grid[i], digits...)
	}
	return grid
}

func FileToIntGrid(file string) ([][]int, error) {
	fs, err := FileToStringSlice(file)
	if err != nil {
		return nil, err
	}

	grid := stringSliceToGrid(fs)
	return grid, nil

}
