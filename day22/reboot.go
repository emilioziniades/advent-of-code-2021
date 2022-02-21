package day22

import (
	"sort"
)

func Reboot(input []string, bounds int) int {
	reactor := make(map[Cuboid]bool)
	cuboids := inputToCuboids(input, bounds)

	/*
		The logic for adding cubes to the reactor map is as follows:
		For every Cuboid from the input:
			If not intersected with any other existing Cuboid in reactor:
				Add to reactor
			Else if there is intersection:
				Split two intersecting cuboids into up to 7 child cuboids, with appropriate on/off status
				Then check if those children cuboids intersect with existing cuboids in reactor:
					If no intersect:
						Add to reactor
					If intersect:
						Recursively do the above procedure

	*/

	var recReboot func(Cuboid)
	recReboot = func(c Cuboid) {
		for r := range reactor {
			//check for overlap
			if isOverlap(c, r) {
				children := Split(r, c)
				//delete overlapping Cuboid since its children will replace it
				delete(reactor, r)
				for _, child := range children {
					recReboot(child)
				}
				return
			}
		}
		// if got here, no overlap, add to reactor
		reactor[c] = c.on
	}

	for len(cuboids) > 0 {
		// pop next Cuboid
		curr := cuboids.PopLeft()
		recReboot(curr)
	}
	return countOn(reactor)
}

// when two cuboids overlap, they can be split into seven distinct cuboids and treated separately
func Split(s1, s2 Cuboid) (children []Cuboid) {

	/*
		This is looking with positive x right, positive y away, and positive z up

		1. Cuboid to the left before intersection
		2. Cuboid on near side before intersection
		3. Cuboid underneath before intersection (always directly below 4)
		4. Cuboid which overlaps with both
		5. Cuboid on top after intersection (always directly above 4)
		6. Cuboid on far side after intersection
		7. Cuboid to the right after intersection

		Parent Split function will always provide a and b such that b's starting point is within a
		But one must also check if a's end point is also within b
	*/

	xS := []int{s1.start.x, s1.end.x, s2.start.x, s2.end.x}
	yS := []int{s1.start.y, s1.end.y, s2.start.y, s2.end.y}
	zS := []int{s1.start.z, s1.end.z, s2.start.z, s2.end.z}
	sort.Ints(xS)
	sort.Ints(yS)
	sort.Ints(zS)

	x1, x2, x3, x4 := xS[0], xS[1], xS[2], xS[3]
	y1, y2, y3, y4 := yS[0], yS[1], yS[2], yS[3]
	z1, z2, z3, z4 := zS[0], zS[1], zS[2], zS[3]

	var c1, c2, c3, c4, c5, c6, c7 *Cuboid
	res := make([]*Cuboid, 0)

	// default state is s1 wholly contains s2 on all sides
	// c4 will always be present, but other six cuboids depend on other factors

	c4 = NewCuboid(x2, y2, z2, x3, y3, z3, s2.on)

	if s1.start.x != s2.start.x {
		c1 = NewCuboid(x1, y1, z1, x2-1, y4, z4, s1.on)
	}
	if s1.start.z != s2.start.z {
		c2 = NewCuboid(x2, y1, z1, x3, y4, z2-1, s1.on)
	}
	if s1.start.y != s2.start.y {
		c3 = NewCuboid(x2, y1, z2, x3, y2-1, z3, s1.on)
	}
	if s1.end.y != s2.end.y {
		c5 = NewCuboid(x2, y3+1, z2, x3, y4, z3, s1.on)
	}
	if s1.end.z != s2.end.z {
		c6 = NewCuboid(x2, y1, z3+1, x3, y4, z4, s1.on)
	}
	if s1.end.x != s2.end.x {
		c7 = NewCuboid(x3+1, y1, z1, x4, y4, z4, s1.on)
	}

	y1y2 := newToggler(y1, y2)
	y3y4 := newToggler(y3, y4)
	z1z2 := newToggler(z1, z2)
	z3z4 := newToggler(z3, z4)

	if s1.start.x > s2.start.x {
		if c1 != nil {
			c1.start.y = y1y2.toggle(c1.start.y)
			c1.start.z = z1z2.toggle(c1.start.z)
			c1.end.y = y3y4.toggle(c1.end.y)
			c1.end.z = z3z4.toggle(c1.end.z)
			c1.on = s2.on
		}
	}

	if s1.start.y > s2.start.y {
		if c1 != nil {
			c1.start.y = y1y2.toggle(c1.start.y)
		}
		if c2 != nil {
			c2.start.y = y1y2.toggle(c2.start.y)
		}
		if c3 != nil {
			c3.on = s2.on
		}
		if c6 != nil {
			c6.start.y = y1y2.toggle(c6.start.y)
		}
		if c7 != nil {
			c7.start.y = y1y2.toggle(c7.start.y)
		}

	}

	if s1.start.z > s2.start.z {
		if c1 != nil {
			c1.start.z = z1z2.toggle(c1.start.z)
		}
		if c2 != nil {
			c2.start.y = y1y2.toggle(c2.start.y)
			c2.end.y = y3y4.toggle(c2.end.y)
			c2.on = s2.on
		}
		if c7 != nil {
			c7.start.z = z1z2.toggle(c7.start.z)
		}

	}

	if s1.end.x < s2.end.x {
		if c7 != nil {
			c7.start.y = y1y2.toggle(c7.start.y)
			c7.start.z = z1z2.toggle(c7.start.z)
			c7.end.y = y3y4.toggle(c7.end.y)
			c7.end.z = z3z4.toggle(c7.end.z)
			c7.on = s2.on
		}
	}

	if s1.end.y < s2.end.y {
		if c1 != nil {
			c1.end.y = y3y4.toggle(c1.end.y)
		}
		if c2 != nil {
			c2.end.y = y3y4.toggle(c2.end.y)
		}
		if c5 != nil {
			c5.on = s2.on
		}
		if c6 != nil {
			c6.end.y = y3y4.toggle(c6.end.y)
		}
		if c7 != nil {
			c7.end.y = y3y4.toggle(c7.end.y)
		}
	}

	if s1.end.z < s2.end.z {
		if c1 != nil {
			c1.end.z = z3z4.toggle(c1.end.z)
		}
		if c6 != nil {
			c6.start.y = y1y2.toggle(c6.start.y)
			c6.end.y = y3y4.toggle(c6.end.y)
			c6.on = s2.on
		}
		if c7 != nil {
			c7.end.z = z3z4.toggle(c7.end.z)
		}

	}

	res = append(res, c1, c2, c3, c4, c5, c6, c7)

	for _, c := range res {
		if c != nil {
			children = append(children, *c)
		}
	}
	return children

}
