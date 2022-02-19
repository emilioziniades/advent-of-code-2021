package day22

type Segment struct {
	start, end int
}

func isAStartinB1(a, b Segment) bool {
	return a.start >= b.start && a.start <= b.end
}

func isAEndinB1(a, b Segment) bool {
	return a.end >= b.start && a.end <= b.end
}

func isAContainsB1(a, b Segment) bool {
	return isAStartinB1(b, a) && isAEndinB1(b, a)
}

func isSegmentOverlap1(a, b Segment) bool {
	return isAStartinB1(a, b) || isAStartinB1(b, a)
}

func Split1D(s1, s2 Segment) (children []Segment) {
	if isAContainsB1(s1, s2) { // s1 wholly contains s2
		return split1D(s1, s2)
	} else if isAContainsB1(s2, s1) { // s2 wholly contains s1
		return split1D(s2, s1)
	} else if isAStartinB1(s2, s1) { // s2 start is in s1
		return split1D(s1, s2)
	} else { // s1 start is in s2
		return split1D(s2, s1)
	}
}

func split1D(s1, s2 Segment) (children []Segment) {
	var x1, x2, x3, x4 int

	x1 = s1.start
	x2 = s2.start

	if isAContainsB1(s1, s2) {
		x3 = s2.end
		x4 = s1.end
	} else {
		x3 = s1.end
		x4 = s2.end
	}

	var c1, c2, c3 *Segment
	res := make([]*Segment, 0)

	c1 = &Segment{x1, x2 - 1}
	c2 = &Segment{x2, x3}
	c3 = &Segment{x3 + 1, x4}

	if s1.start == s2.start {
		c1 = nil
	}

	if s1.end == s2.end {
		c3 = nil
	}
	res = append(res, c1, c2, c3)

	for _, c := range res {
		if c != nil {
			children = append(children, *c)
		}
	}
	return children

}
