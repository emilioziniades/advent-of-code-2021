package day24_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day24"
)

func TestRemakeInstructions(t *testing.T) {
	got := day24.RemakeInstructions()

	wantByte, _ := os.ReadFile("input.txt")
	want := string(wantByte)

	if got != want {
		msg := fmt.Sprintf("\nGOT:\n\n%#v\n\nWANT:\n\n%#v", got, want)
		t.Error(msg)
	}
}

// ensures that our step function stays in line with raw program
func TestStepAll(t *testing.T) {
	program := day24.LoadProgram("input.txt")
	for n := day24.MaxInput; n >= day24.MinInput; n -= 123325798721 {

		want := day24.Run(program, n, "z")
		got := day24.StepAll(n)
		if got != want {
			t.Errorf("TestStepAll: got %v, want %v for %v", got, want, n)
		}

	}
}
