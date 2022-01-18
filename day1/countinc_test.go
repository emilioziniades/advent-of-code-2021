package day1_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day1"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/1/input", "1.in")
	if err != nil {
		log.Fatal(err)
	}
}
func TestCountInc(t *testing.T) {
	testCountInc(t, "1.ex", 7)
	testCountInc(t, "1.in", 1529)
}

func TestCountIncThree(t *testing.T) {
	testCountIncThree(t, "1.ex", 5)
	testCountIncThree(t, "1.in", 1567)
}

func testCountInc(t *testing.T, file string, want int) {
	in, err := parse.FileToIntSlice(file)
	if err != nil {
		log.Fatalf("countinc: file read error: %s", err)
	}
	count := day1.CountInc(in)
	if count != want {
		t.Fatalf("countinc: wanted %d, got %d", want, count)
	}
	fmt.Printf("Got %d, wanted %d for %s\n", count, want, file)
}

func testCountIncThree(t *testing.T, file string, want int) {
	in, err := parse.FileToIntSlice(file)
	if err != nil {
		log.Fatalf("countincthree: file read error: %s", err)
	}
	count := day1.CountIncThree(in)
	if count != want {
		t.Fatalf("countincthree: wanted %d, got %d", want, count)
	}
	fmt.Printf("Got %d, wanted %d for %s\n", count, want, file)
}

func parseInput(file string) ([]int, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]int, 0)
	input := bufio.NewScanner(f)
	for input.Scan() {
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, n)
	}
	return res, nil

}
