package day22

import (
	"sort"
)

type point2 struct {
	x, y int
}

type Square struct {
	start, end point2
}

func Split2D(s1, s2 Square) (children []Square) {

	xS := []int{s1.start.x, s1.end.x, s2.start.x, s2.end.x}
	yS := []int{s1.start.y, s1.end.y, s2.start.y, s2.end.y}
	sort.Ints(xS)
	sort.Ints(yS)

	x1, x2, x3, x4 := xS[0], xS[1], xS[2], xS[3]
	y1, y2, y3, y4 := yS[0], yS[1], yS[2], yS[3]

	var c1, c2, c3, c4, c5 *Square
	res := make([]*Square, 0)

	// default state is s1 wholly contains s2 on all sides
	c1 = &Square{point2{x1, y1}, point2{x2 - 1, y4}}
	c2 = &Square{point2{x2, y1}, point2{x3, y2 - 1}}
	c3 = &Square{point2{x2, y2}, point2{x3, y3}}
	c4 = &Square{point2{x2, y3 + 1}, point2{x3, y4}}
	c5 = &Square{point2{x3 + 1, y1}, point2{x4, y4}}

	y1y2 := newToggler(y1, y2)
	y3y4 := newToggler(y3, y4)

	if s1.start.x > s2.start.x {
		c1.start.y = y1y2.toggle(c1.start.y)
		c1.end.y = y3y4.toggle(c1.end.y)
	}

	if s1.start.y > s2.start.y {
		c1.start.y = y1y2.toggle(c1.start.y)
		c5.start.y = y1y2.toggle(c5.start.y)
	}

	if s1.end.x < s2.end.x {
		c5.start.y = y1y2.toggle(c5.start.y)
		c5.end.y = y3y4.toggle(c5.end.y)
	}

	if s1.end.y < s2.end.y {
		c1.end.y = y3y4.toggle(c1.end.y)
		c5.end.y = y3y4.toggle(c5.end.y)
	}

	if s1.start.x == s2.start.x {
		c1 = nil
	}
	if s1.start.y == s2.start.y {
		c2 = nil
	}
	if s1.end.y == s2.end.y {
		c4 = nil
	}
	if s1.end.x == s2.end.x {
		c5 = nil
	}

	res = append(res, c1, c2, c3, c4, c5)

	for _, c := range res {
		if c != nil {
			children = append(children, *c)
		}
	}
	return children

}
