package day18_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/emilioziniades/adventofcode2021/day18"
	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/18/input", "18.in")
	if err != nil {
		log.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	test := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}
	for _, e := range test {
		_ = day18.Parse(e)
	}
}

func TestAdd(t *testing.T) {
	ex, _ := parse.FileToStringSlice("18.ex")
	tests := []struct {
		in   []string
		want string
	}{
		{[]string{"[1,2]", "[[3,4],5]"}, "[[1,2],[[3,4],5]]"},
		{[]string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]"}, "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]"}, "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{[]string{"[1,1]", "[2,2]", "[3,3]", "[4,4]", "[5,5]", "[6,6]"}, "[[[[5,0],[7,4]],[5,5]],[6,6]]"},
		{ex, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
	}

	for _, tt := range tests {
		fmt.Println("testing: ", tt.in)
		got := day18.Add(tt.in)
		format := "\ngot:\n\t%s\nwant:\n\t%s\nfor:\n\t%v\n"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.in)
		} else {
			t.Logf(format, got, tt.want, tt.in)
		}
	}
}

func TestExplode(t *testing.T) {
	test := []string{
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", // should not explode
	}

	for _, tt := range test {
		day18.Explode(tt)
	}
}

func TestSplit(t *testing.T) {
	test := []string{
		"[10,0]",
		"[11,0]",
		"[[[[[9,8],1],2],3],4]", // should not split
	}

	for _, tt := range test {
		day18.Split(tt)
	}
}

func TestMagnitude(t *testing.T) {
	ex2, _ := parse.FileToStringSlice("18.ex2")
	in, _ := parse.FileToStringSlice("18.in")
	tests := []struct {
		in   []string
		want int
	}{
		{[]string{"[9,1]"}, 29},
		{[]string{"[[9,1],[1,9]]"}, 129},
		{[]string{"[[1,2],[[3,4],5]]"}, 143},
		{[]string{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"}, 1384},
		{[]string{"[[[[1,1],[2,2]],[3,3]],[4,4]]"}, 445},
		{[]string{"[[[[3,0],[5,3]],[4,4]],[5,5]]"}, 791},
		{[]string{"[[[[5,0],[7,4]],[5,5]],[6,6]]"}, 1137},
		{[]string{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}, 3488},
		{ex2, 4140},
		{in, 4347},
	}
	for _, tt := range tests {
		got := day18.Magnitude(tt.in)
		format := "got %d, wanted %d, for %s"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.in)
		} else {
			t.Logf(format, got, tt.want, tt.in)
		}
	}
}

func TestMaxMagnitude(t *testing.T) {
	tests := []struct {
		file string
		want int
	}{
		{"18.ex2", 3993},
		{"18.in", 4721},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			log.Fatal(err)
		}
		got := day18.MaxMagnitude(in)
		format := "got %d, wanted %d, for %s"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.file)
		} else {
			t.Logf(format, got, tt.want, tt.file)
		}

	}
}
