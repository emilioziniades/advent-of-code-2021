package day2_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day2"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/2/input", "day2-input.txt")
	if err != nil {
		log.Fatal(err)
	}
}
func TestMove(t *testing.T) {
	testMove(t, "day2-example.txt", 150)
	testMove(t, "day2-input.txt", 1746616)
}

func testMove(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		log.Fatalf("move: file read error: %s", err)
	}

	prod := day2.Move(in)
	if prod != want {
		t.Fatalf("Move: wanted %d, got %d\n", want, prod)
	}
	fmt.Printf("Move: got %d, wanted %d for %s\n", prod, want, file)
}

func TestMoveAim(t *testing.T) {
	testMoveAim(t, "day2-example.txt", 900)
	testMoveAim(t, "day2-input.txt", 1741971043)
}

func testMoveAim(t *testing.T, file string, want int) {
	in, err := parse.FileToStringSlice(file)
	if err != nil {
		log.Fatalf("MoveAim: file read error: %s", err)
	}

	prod := day2.MoveAim(in)
	if prod != want {
		t.Fatalf("MoveAim: wanted %d, got %d\n", want, prod)
	}
	fmt.Printf("MoveAim: got %d, wanted %d for %s\n", prod, want, file)
}
