package day24_test

import (
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day24"
	"github.com/emilioziniades/adventofcode2021/fetch"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/24/input", "input.txt")
	if err != nil {
		log.Fatal(err)
	}
}
func TestPrograms(t *testing.T) {

	type testcase struct {
		outvar string
		input  int
		want   int
	}

	tests := []struct {
		filename  string
		testcases []testcase
	}{
		{
			"example1.txt",
			[]testcase{
				{
					"x",
					1,
					-1,
				},
				{
					"x",
					9,
					-9,
				},
			},
		},
		{
			"example2.txt",
			[]testcase{
				{
					"z",
					39,
					1,
				},
				{
					"z",
					36,
					0,
				},
			},
		},
		{
			"example3.txt",
			[]testcase{
				{
					"w",
					7,
					0,
				},
				{
					"x",
					7,
					1,
				},
				{
					"y",
					7,
					1,
				},
				{
					"z",
					7,
					1,
				},
				{
					"w",
					9,
					1,
				},
				{
					"x",
					9,
					0,
				},
				{
					"y",
					9,
					0,
				},
				{
					"z",
					9,
					1,
				},
			},
		},
	}

	for _, tt := range tests {
		program := day24.LoadProgram(tt.filename)
		for _, testcase := range tt.testcases {
			got := day24.Run(program, testcase.input, testcase.outvar)
			if got != testcase.want {
				t.Errorf(
					"TestPrograms: got %v wanted %v for program: \n %#v \n inputs: %v",
					got,
					testcase.want,
					program,
					testcase.input,
				)
			}

		}
	}
}

func testValidateModelNumber(t *testing.T) {
	tests := []struct {
		filename string
		want     int
	}{
		{
			"input.txt",
			0,
		},
	}

	for _, tt := range tests {
		if got := day24.ValidateModelNumber(tt.filename); got != tt.want {
			t.Errorf("TestValidateModelNumber: got %v, wanted %v for %v", got, tt.want, tt.filename)
		}
	}

}
