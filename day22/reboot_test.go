package day22

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/emilioziniades/adventofcode2021/fetch"
	"github.com/emilioziniades/adventofcode2021/parse"
)

func init() {
	err := fetch.Data("https://adventofcode.com/2021/day/22/input", "22.in")
	if err != nil {
		panic(err)
	}
}

func TestReboot(t *testing.T) {
	var tests = []struct {
		file  string
		want  int
		limit int
	}{
		// {"22.me", 4, 50},
		// {"22.si", 39, 50},
		{"22.ex", 590784, 50},
		// {"22.ex2", 474140, 50},
		// {"22.in", 553201, 50},
		// {"22.ex2", 474140, math.MaxInt},
	}
	for _, tt := range tests {
		in, err := parse.FileToStringSlice(tt.file)
		if err != nil {
			panic(err)
		}

		got := Reboot(in, tt.limit)
		format := "got %d, wanted %d for %s"
		if got != tt.want {
			t.Fatalf(format, got, tt.want, tt.file)
		} else {
			t.Logf(format, got, tt.want, tt.file)
		}
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		c1, c2 Cuboid
		want   int
		title  string
	}{
		{
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			Cuboid{Point{10, 10, 10}, Point{10, 10, 10}, false},
			46,
			"testing on/off",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{10, 10, 10}, false},
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			46,
			"testing on/off in reverse",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			Cuboid{Point{11, 11, 11}, Point{13, 13, 13}, true},
			46,
			"A. s1.e.x < s2.e.y & s1.e.y < s2.e.y & s1.e.z < s2.e.z",
		},
		{
			Cuboid{Point{11, 11, 11}, Point{13, 13, 13}, true},
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			46,
			"A reversed. s1.s.x > s2.s.y & s1.s.y > s2.s.y & s1.s.z > s2.s.z ",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{13, 13, 13}, true},
			Cuboid{Point{12, 12, 12}, Point{15, 15, 15}, true},
			120,
			"",
		},
		{
			Cuboid{Point{12, 12, 12}, Point{15, 15, 15}, true},
			Cuboid{Point{10, 10, 10}, Point{13, 13, 13}, true},
			120,
			"",
		},
		{
			Cuboid{Point{9, 9, 9}, Point{11, 11, 11}, true},
			Cuboid{Point{10, 10, 10}, Point{10, 12, 12}, true},
			32,
			"C. s2.end.y > s1.end.y and s1.end.z < s2.end.z",
		},
		{
			Cuboid{Point{11, 9, 9}, Point{11, 11, 11}, true},
			Cuboid{Point{11, 10, 10}, Point{12, 10, 12}, true},
			13,
			"C reversed",
		},
		{
			Cuboid{Point{11, 10, 10}, Point{12, 10, 12}, true},
			Cuboid{Point{11, 9, 9}, Point{11, 11, 11}, true},
			13,
			"Funny shape",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{10, 10, 10}, true},
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			27,
			"E. wholly contained, but shared starting x, y and z",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			Cuboid{Point{10, 10, 10}, Point{10, 10, 10}, true},
			27,
			"E reversed.",
		},
		{
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			Cuboid{Point{11, 11, 11}, Point{13, 13, 11}, true},
			30,
			"F. end x and end y",
		},
		{
			Cuboid{Point{11, 11, 11}, Point{13, 13, 11}, true},
			Cuboid{Point{10, 10, 10}, Point{12, 12, 12}, true},
			30,
			"F reversed",
		},
	}
	for _, tt := range tests {
		children := Split(tt.c1, tt.c2)
		count := 0
		for _, c := range children {
			count += c.Volume()
		}
		fmt.Println(children)
		format := "wanted %d, got %d, for Cuboids %v and %v (%s)"
		if count != tt.want {
			t.Errorf(format, tt.want, count, tt.c1, tt.c2, tt.title)
		} else {
			t.Logf(format, tt.want, count, tt.c1, tt.c2, tt.title)
		}
	}
}

func Test1DAnd2D(t *testing.T) {
	test1DSplit(t)
	test2DSplit(t)
}

func test1DSplit(t *testing.T) {
	var tests = []struct {
		a, b  Segment
		want  []Segment
		title string
	}{
		{Segment{1, 4}, Segment{4, 5}, []Segment{{1, 3}, {4, 4}, {5, 5}}, "case 1"},                   // case 1
		{Segment{1, 4}, Segment{3, 4}, []Segment{{1, 2}, {3, 4}}, "case 2"},                           // case 2
		{Segment{1, 4}, Segment{2, 3}, []Segment{{1, 1}, {2, 3}, {4, 4}}, "case 3"},                   // case 3
		{Segment{1, 4}, Segment{1, 2}, []Segment{{1, 2}, {3, 4}}, "case 4"},                           // case 4
		{Segment{1, 4}, Segment{0, 1}, []Segment{{0, 0}, {1, 1}, {2, 4}}, "case 5"},                   // case 5
		{Segment{1, 2}, Segment{1, 4}, []Segment{{1, 2}, {3, 4}}, "case 4b"},                          // case 4b
		{Segment{3, 4}, Segment{1, 4}, []Segment{{1, 2}, {3, 4}}, "case 2b"},                          // case 2b
		{Segment{2, 3}, Segment{1, 4}, []Segment{{1, 1}, {2, 3}, {4, 4}}, "case 3b"},                  // case 3b
		{Segment{33, 67}, Segment{50, 105}, []Segment{{33, 49}, {50, 67}, {68, 105}}, "case 1, hard"}, // case 1
		{Segment{11, 24}, Segment{5, 30}, []Segment{{5, 10}, {11, 24}, {25, 30}}, "case 3b, hard"},
		{Segment{25, 50}, Segment{25, 80}, []Segment{{25, 50}, {51, 80}}, "case 4b, hard"},
		{Segment{99, 99}, Segment{99, 99}, []Segment{{99, 99}}, "random case"},
	}
	for _, tt := range tests {
		got := Split1D(tt.a, tt.b)
		format := "%s: got %v, wanted %v, for %v and %v\n"
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("FAIL "+format, tt.title, got, tt.want, tt.a, tt.b)
		} else {
			t.Logf("PASS "+format, tt.title, got, tt.want, tt.a, tt.b)
		}

	}
}

func test2DSplit(t *testing.T) {

	var tests = []struct {
		a, b  Square
		want  []Square
		title string
	}{
		{
			Square{point2{1, 1}, point2{3, 3}},
			Square{point2{2, 2}, point2{4, 4}},
			[]Square{{point2{1, 1}, point2{1, 3}}, {point2{2, 1}, point2{3, 1}}, {point2{2, 2}, point2{3, 3}}, {point2{2, 4}, point2{3, 4}}, {point2{4, 2}, point2{4, 4}}},
			"normal overlap",
		},
		{
			Square{point2{2, 2}, point2{4, 4}},
			Square{point2{1, 1}, point2{3, 3}},
			[]Square{{point2{1, 1}, point2{1, 3}}, {point2{2, 1}, point2{3, 1}}, {point2{2, 2}, point2{3, 3}}, {point2{2, 4}, point2{3, 4}}, {point2{4, 2}, point2{4, 4}}},
			"normal overlap with order reversed",
		},
		{
			Square{point2{1, 1}, point2{4, 4}},
			Square{point2{3, 3}, point2{3, 3}},
			[]Square{{point2{1, 1}, point2{2, 4}}, {point2{3, 1}, point2{3, 2}}, {point2{3, 3}, point2{3, 3}}, {point2{3, 4}, point2{3, 4}}, {point2{4, 1}, point2{4, 4}}},
			"x, y both fully contained, no sides touching",
		},
		{
			Square{point2{1, 1}, point2{3, 3}},
			Square{point2{2, 1}, point2{2, 3}},
			[]Square{{point2{1, 1}, point2{1, 3}}, {point2{2, 1}, point2{2, 3}}, {point2{3, 1}, point2{3, 3}}},
			"x, y both contained, top, bottom touching",
		},
		{
			Square{point2{1, 1}, point2{3, 3}},
			Square{point2{2, 2}, point2{2, 4}},
			[]Square{{point2{1, 1}, point2{1, 3}}, {point2{2, 1}, point2{2, 1}}, {point2{2, 2}, point2{2, 3}}, {point2{2, 4}, point2{2, 4}}, {point2{3, 1}, point2{3, 3}}},
			"x fully contained, y normal",
		},
		{
			Square{point2{1, 1}, point2{3, 3}},
			Square{point2{2, 2}, point2{2, 4}},
			[]Square{{point2{1, 1}, point2{1, 3}}, {point2{2, 1}, point2{2, 1}}, {point2{2, 2}, point2{2, 3}}, {point2{2, 4}, point2{2, 4}}, {point2{3, 1}, point2{3, 3}}},
			"x fully contained, y normal",
		},
		{
			Square{point2{1, 1}, point2{3, 3}},
			Square{point2{1, 1}, point2{2, 3}},
			[]Square{{point2{1, 1}, point2{2, 3}}, {point2{3, 1}, point2{3, 3}}},
			"x, y both contained, left, right top bottom touching",
		},
		{
			Square{point2{0, 1}, point2{3, 3}},
			Square{point2{1, 0}, point2{2, 2}},
			[]Square{{point2{0, 1}, point2{0, 3}}, {point2{1, 0}, point2{2, 0}}, {point2{1, 1}, point2{2, 2}}, {point2{1, 3}, point2{2, 3}}, {point2{3, 1}, point2{3, 3}}},
			"x, y both contained, left, right top bottom touching",
		},
	}
	for _, tt := range tests {
		got := Split2D(tt.a, tt.b)
		format := "%s: \n\tgot  %v, \n\twant %v, for %v and %v\n"
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("FAIL "+format, tt.title, got, tt.want, tt.a, tt.b)
		} else {
			t.Logf("PASS "+format, tt.title, got, tt.want, tt.a, tt.b)
		}

	}
}
