package fetch

import (
	"bufio"
	"os"
	"strconv"
)

func ParseInputString(file string) ([]string, error) {
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

func ParseInputInt(file string) ([]int, error) {
	stringRes, err := ParseInputString(file)
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
