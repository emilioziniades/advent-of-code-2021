package main

import "fmt"

type point struct {
	x, y int
}

type rectangle struct {
	start, end point
}

func main() {
	a := point{10, 10}
	b := point{11, 11}
	c := point{12, 12}
	d := point{13, 13}
	r1 := rectangle{a, c}
	r2 := rectangle{b, d}
	fmt.Println(r1, r2)
	splits := split(r1, r2)
	fmt.Println(splits)
}

func split(r1, r2 rectangle) []rectangle {
	res := make([]rectangle, 0)
	nr1 := rectangle{r1.start, point{r2.start.x, r1.end.y}}
	nr2 := rectangle{point{r2.start.x, r1.start.y}, point{r1.end.x, r2.start.y}}
	nr3 := rectangle{r2.start, r1.end}
	nr4 := rectangle{point{r2.start.x, r1.end.y}, point{r1.end.x, r2.end.y}}
	nr5 := rectangle{point{r1.end.x, r2.start.y}, r2.end}
	res = append(res, nr1, nr2, nr3, nr4, nr5)
	return res
}
