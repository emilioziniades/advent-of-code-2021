package day22

import "log"

type Point struct {
	x, y, z int
}

type Cuboid struct {
	start, end Point
	on         bool
}

func NewCuboid(xS, yS, zS, xE, yE, zE int, on bool) *Cuboid {
	format := "NewCuboid: start %s greater than end %s. %v"
	c := &Cuboid{Point{xS, yS, zS}, Point{xE, yE, zE}, on}
	if xS > xE {
		log.Fatalf(format, "x", "x", *c)
	}
	if yS > yE {
		log.Fatalf(format, "y", "y", *c)
	}
	if zS > zE {
		log.Fatalf(format, "z", "z", *c)
	}
	return c
}

func (c Cuboid) Volume() int {
	return (c.end.x - c.start.x + 1) * (c.end.y - c.start.y + 1) * (c.end.z - c.start.z + 1)
}

type toggler struct {
	x1, x2, toggleVal int
}

func newToggler(x1, x2 int) toggler {
	return toggler{
		toggleVal: x1 ^ x2,
		x1:        x1,
		x2:        x2,
	}
}

func (t toggler) toggle(x int) int {
	if x != t.x1 && x != t.x2 {
		panic("toggle: can't toggle value not included in toggleVal")
	}
	return t.toggleVal ^ x
}
