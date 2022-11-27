package day24_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day24"
)

func TestRemakeInstructions(t *testing.T) {
	filename := "input.txt"
	changingValues := day24.GetChangingValues(filename)
	got := day24.RemakeInstructions(changingValues)

	wantByte, _ := os.ReadFile(filename)
	want := string(wantByte)

	if got != want {
		msg := fmt.Sprintf("\nGOT:\n\n%#v\n\nWANT:\n\n%#v", got, want)
		t.Error(msg)
	}
}

// ensures that our step function stays in line with raw program
func TestStepAll(t *testing.T) {
	filename := "input.txt"
	program := day24.LoadProgram(filename)
	changingValues := day24.GetChangingValues(filename)
	for n := day24.MaxInput; n >= day24.MinInput; n -= 123325798721 {

		want := day24.Run(program, n, "z")
		got := day24.StepAll(n, changingValues)
		if got != want {
			t.Errorf("TestStepAll: got %v, want %v for %v", got, want, n)
		}

	}
}
