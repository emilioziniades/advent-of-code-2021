package day18

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"text/scanner"
)

var depth int

type tree struct {
	value       int
	num         bool
	left, right *tree
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next() { lex.token = lex.scan.Scan() }

func (lex *lexer) text() string { return lex.scan.TokenText() }

//  expr = num
//		 | '[' expr ',' expr ']'

func Parse(s string) *tree {
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(s))
	lex.next()

	root := &tree{}
	parse(lex, root)
	//	 printTree(*root)
	return root
}

func parse(lex *lexer, t *tree) {
	switch lex.token {
	case '[':
		lex.next()
		t.left = &tree{value: 0, num: false}
		parse(lex, t.left)

		lex.next() //consume ,

		lex.next()
		t.right = &tree{value: 0, num: false}
		parse(lex, t.right)
	case scanner.Int:
		n, err := strconv.Atoi(lex.text())
		if err != nil {
			log.Fatal(err)
		}
		t.value, t.num = n, true
	case ',':
		lex.next()
		parse(lex, t)
	case ']':
		lex.next()
		parse(lex, t)
	}
}

func Add(ss []string) string {
	trees := make([]*tree, 0)
	for _, s := range ss {
		trees = append(trees, Parse(s))
	}

	overall := trees[0]
	for i := 1; i < len(trees); i++ {
		overall = add(overall, trees[i])
		var curr *tree
		ok := true
		for ok {
			curr, ok = isExplode(overall, 0)
			list := collapse(overall)
			if ok {
				explode(curr, list)
				continue
			}

			curr, ok = isSplit(overall)
			if ok {
				split(curr)
				continue
			}
		}
		//		printTree(*overall)
	}

	return printRes(*overall)
}

func add(t1, t2 *tree) *tree {
	root := new(tree)
	root.left = t1
	root.right = t2
	return root
}

func Explode(s string) *tree {
	t := Parse(s)
	curr, ok := isExplode(t, 0)
	treeList := collapse(t)

	if ok {
		explode(curr, treeList)
	}
	//	printTree(*t)
	//	fmt.Println(collapse(t))
	return t
}

func explode(curr *tree, treeList []*tree) {
	nL := pos(curr.left, treeList)
	nR := pos(curr.right, treeList)

	// Steps to explode

	//1. pair's left value added to first number left of pair (if any)
	if nL > 0 {
		// has value to the left
		treeList[nL-1].value += curr.left.value
	}

	//2. pair's right value added to first number right of pair (if any)
	if nR < len(treeList)-1 {
		// has value to the right
		treeList[nR+1].value += curr.right.value
	}

	//3. entire pair replaced with number 0
	curr.value = 0
	curr.num = true
	curr.left, curr.right = nil, nil

}

func isExplode(t *tree, depth int) (*tree, bool) {
	if depth >= 4 && t.left != nil && t.right != nil {
		//		fmt.Println("should explode: ", t.left.value, t.right.value)
		return t, true
	}
	if t.left == nil && t.right == nil {
		return nil, false
	}
	var currTree *tree
	var ok bool

	if t.left != nil {
		currTree, ok = isExplode(t.left, depth+1)
	}
	if !ok {
		currTree, ok = isExplode(t.right, depth+1)
	}
	return currTree, ok
}

func Split(s string) *tree {
	t := Parse(s)
	curr, ok := isSplit(t)
	fmt.Println(curr, ok)
	if ok {
		split(curr)
	}
	//	printTree(*t)
	return t
}

func split(curr *tree) {
	n := float64(curr.value)
	floor := int(math.Floor(n / 2))
	ceil := int(math.Ceil(n / 2))
	//	fmt.Println(n, floor, ceil)

	curr.value = 0
	curr.num = false
	curr.left = &tree{value: floor, num: true}
	curr.right = &tree{value: ceil, num: true}

}

func isSplit(t *tree) (*tree, bool) {
	if t.value > 9 {
		//		fmt.Println("should split: ", t.value)
		return t, true
	}

	if t.left == nil && t.right == nil {
		return nil, false
	}

	var curr *tree
	var ok bool

	if t.left != nil {
		curr, ok = isSplit(t.left)
	}

	if !ok {
		curr, ok = isSplit(t.right)
	}

	return curr, ok
}

func Magnitude(ss []string) int {
	final := Parse(Add(ss))

	var recMagnitude func(*tree) int
	recMagnitude = func(t *tree) int {
		if t.num {
			return t.value
		} else {
			return 3*recMagnitude(t.left) + 2*recMagnitude(t.right)
		}
	}
	return recMagnitude(final)
}

func MaxMagnitude(ss []string) int {
	max := math.MinInt
	n := len(ss) - 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			curr := make([]string, 0)
			curr = append(curr, ss[i], ss[j])
			mag := Magnitude(curr)
			if mag > max {
				max = mag
			}
		}
	}
	return max
}

func pos(t *tree, tl []*tree) int {
	for i, e := range tl {
		if t == e {
			return i
		}
	}
	return -1
}

func collapse(t *tree) []*tree {
	res := make([]*tree, 0)
	var recCollapse func(*tree)
	recCollapse = func(t *tree) {
		if !t.num {
			recCollapse(t.left)
			recCollapse(t.right)
		}
		if t.num {
			res = append(res, t)
		}
	}
	recCollapse(t)
	return res
}

func printList(tl []*tree) {
	for _, e := range tl {
		fmt.Printf("%d ", (*e).value)
	}
	fmt.Println("")
}

func printTree(t tree) {
	if t.num {
		fmt.Printf("%*s%d\n", depth*4, "", t.value)
	} else {
		fmt.Printf("%*s.\n", depth*4, "")
	}
	depth++
	if t.left != nil {
		printTree(*t.left)
	}
	if t.right != nil {
		printTree(*t.right)
	}
	depth--
}

func printRes(t tree) string {
	var b strings.Builder
	var recPrint func(tree)
	recPrint = func(t tree) {
		fmt.Fprint(&b, "[")
		if t.left.num {
			fmt.Fprint(&b, t.left.value)
		} else {
			recPrint(*t.left)
		}
		fmt.Fprint(&b, ",")
		if t.right.num {
			fmt.Fprint(&b, t.right.value)
		} else {
			recPrint(*t.right)
		}
		fmt.Fprint(&b, "]")
	}
	recPrint(t)
	return b.String()
}
