package day16

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

type item struct {
	value int
	op    bool
}

type tree struct {
	it       item
	children []*tree
}

var opMap = map[int]string{
	0: "sum",
	1: "product",
	2: "minimum",
	3: "maximum",
	4: "literal",
	5: "greater",
	6: "less",
	7: "equal",
}

var versionCount int
var depth int

func Bits(file []string) (int, int) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("")
	if len(file) != 1 {
		log.Fatal("day16: file should be one line")
	}

	hex := file[0]
	dec := new(big.Int)
	if _, ok := dec.SetString(hex, 16); !ok {
		log.Fatalf("Bits: converting %s", hex)
	}
	//pads with leading zeroes so that num bits = 4 * hex digits
	bin := fmt.Sprintf("%0*b", len(hex)*4, dec)
	fmt.Println(bin, hex)

	versionCount = 0
	r := strings.NewReader(bin)
	t := new(tree)
	parsePacket(r, t, 100, -1)
	t = t.children[0]
	printTree(*t)
	fmt.Println("")
	return versionCount, reduceTree(t)
}

func parsePacket(r *strings.Reader, t *tree, length, lengthType int) {
	nRead := 0
	for nRead < length {
		version := readN(r, 3)
		typeId := readN(r, 3)
		versionCount += version
		bitsRead := 6
		switch typeId {
		case 4:
			value := new(strings.Builder)
			for {
				done := readN(r, 1)
				vals := readNStr(r, 4)
				value.WriteString(vals)
				bitsRead += 5
				if done == 0 {
					break
				}
			}
			v := parseBits(value.String())
			child := tree{item{v, false}, make([]*tree, 0)}
			t.children = append(t.children, &child)
		default:
			lengthTypeId := readN(r, 1)
			bitsRead += 1
			var l int
			switch lengthTypeId {
			case 1:
				l = readN(r, 11)
				bitsRead += 11
			case 0:
				l = readN(r, 15)
				bitsRead += 15
			}

			child := tree{item{typeId, true}, make([]*tree, 0)}
			t.children = append(t.children, &child)

			oldLen := r.Len()
			parsePacket(r, &child, l, lengthTypeId)
			bitsRead += oldLen - r.Len()
		}
		switch lengthType {
		case 1:
			// length is number of subpackets
			nRead += 1
		case 0:
			// length is total length of subpackets in bits
			nRead += bitsRead
		case -1:
			// initial packet
			nRead = length

		}
	}
}

func reduceTree(t *tree) int {
	curr := t.it
	if curr.op {
		switch curr.value {
		case 0:
			sum := 0
			for _, e := range t.children {
				sum += reduceTree(e)
			}
			return sum
		case 1:
			product := 1
			for _, e := range t.children {
				product *= reduceTree(e)
			}
			return product
		case 2:
			min := math.MaxInt
			for _, e := range t.children {
				if r := reduceTree(e); r < min {
					min = r
				}
			}
			return min
		case 3:
			max := 0
			for _, e := range t.children {
				if r := reduceTree(e); r > max {
					max = r
				}
			}
			return max
		case 5:
			greater := 0
			first, second := reduceTree(t.children[0]), reduceTree(t.children[1])
			if first > second {
				greater = 1
			}
			return greater
		case 6:
			less := 0
			first, second := reduceTree(t.children[0]), reduceTree(t.children[1])
			if first < second {
				less = 1
			}
			return less
		case 7:
			equal := 0
			first, second := reduceTree(t.children[0]), reduceTree(t.children[1])
			if first == second {
				equal = 1
			}
			return equal
		default:
			return 10000000
		}
	} else {
		return curr.value
	}
}

func printTree(t tree) {
	curr := t.it
	if curr.op {
		fmt.Printf("%*s%s\n", depth*4, "", opMap[curr.value])
	} else {
		fmt.Printf("%*s%d\n", depth*4, "", curr.value)
	}
	depth++
	for _, e := range t.children {
		printTree(*e)
	}
	depth--
}

//readN reads n runes from r and returns it as an int
func readN(r *strings.Reader, n int) int {
	s := new(strings.Builder)
	for i := 0; i < n; i++ {
		ch, _, err := r.ReadRune()
		if err != nil {
			//			debug.PrintStack()
			log.Fatal(err)
		}
		s.WriteRune(ch)
	}
	return parseBits(s.String())
}

//readNStr reads n runes from r and returns it as an string
func readNStr(r *strings.Reader, n int) string {
	s := new(strings.Builder)
	for i := 0; i < n; i++ {
		ch, _, err := r.ReadRune()
		if err != nil {
			//			debug.PrintStack()
			log.Fatal(err)
		}
		s.WriteRune(ch)
	}
	return s.String()
}

func parseBits(s string) int {
	i, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
