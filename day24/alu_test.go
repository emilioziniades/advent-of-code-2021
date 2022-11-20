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
		inputs []int
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
					[]int{1},
					-1,
				},
				{
					"x",
					[]int{42},
					-42,
				},
			},
		},
		{
			"example2.txt",
			[]testcase{
				{
					"z",
					[]int{3, 9},
					1,
				},
				{
					"z",
					[]int{3, 6},
					0,
				},
			},
		},
		{
			"example3.txt",
			[]testcase{
				{
					"w",
					[]int{15},
					1,
				},
				{
					"x",
					[]int{15},
					1,
				},
				{
					"y",
					[]int{15},
					1,
				},
				{
					"z",
					[]int{15},
					1,
				},
				{
					"w",
					[]int{42},
					1,
				},
				{
					"x",
					[]int{42},
					0,
				},
				{
					"y",
					[]int{42},
					1,
				},
				{
					"z",
					[]int{42},
					0,
				},
			},
		},
	}

	for _, tt := range tests {
		program := day24.LoadProgram(tt.filename)
		for _, testcase := range tt.testcases {
			got := day24.Run(program, testcase.inputs, testcase.outvar)
			if got != testcase.want {
				t.Errorf(
					"TestPrograms: got %v wanted %v for program: \n %#v \n inputs: %v",
					got,
					testcase.want,
					program,
					testcase.inputs,
				)
			}

		}
	}
}

func TestValidateModelNumber(t *testing.T) {
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
