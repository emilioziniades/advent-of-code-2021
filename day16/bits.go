package day16

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/emilioziniades/adventofcode2021/stack"
)

type item struct {
	value int
	op    bool
}

var opMap = map[int]string{
	0: "sum",
	1: "product",
	2: "minimum",
	3: "maximum",
	4: "literal",
	5: "less",
	6: "greater",
	7: "equal",
}

type tree struct {
	it item
	children []*tree
}

var versionCount int
var depth int

var s stack.Stack[item]

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
	//pads with leading zeroes so that there are exactly 4 times as many bits as hex digits
	bin := fmt.Sprintf("%0*b", len(hex)*4, dec)
	fmt.Println(bin, hex)

	s = stack.New[item]()
	r := strings.NewReader(bin)
	versionCount = 0
//	t := new(tree)
	parsePacket(r, 100, -1)
	//	fmt.Println("\nversion count: ", versionCount)
	fmt.Println(s)
	return versionCount, parseStack(s)
}

func parsePacket(r *strings.Reader, length, lengthType int) {
	depth++
	//	fmt.Printf("%*snew call to parsePacket with length %d, type ID %d and %d remaining bits\n", depth*4, "", length, lengthType, r.Len())
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
			//			fmt.Printf("%*sliteral packet -> version: %d, type ID: %d, literal value: %d\n", depth*4, "", version, typeId, parseBits(value.String()))
			v := parseBits(value.String())
			s.Push(item{v, false})
			fmt.Printf("%*sliteral -> %d\n", depth*4, "", v)
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

			//			fmt.Printf("%*soperator packet -> version: %d, type ID: %d, length type ID: %d, length info: %d\n", depth*4, "", version, typeId, lengthTypeId, l)
			fmt.Printf("%*soperator ->  %s\n", depth*4, "", opMap[typeId])
			s.Push(item{typeId, true})
			oldLen := r.Len()
			parsePacket(r, l, lengthTypeId)
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

		//		fmt.Printf("%*sremaining bits: %d\n", depth*4, "", r.Len()
		//		fmt.Printf("%*slength: %d, nRead: %d\n", depth*4, "", length, nRead)
	}
	depth--
}

func parseStack(s stack.Stack[item]) int {
	valStack := stack.New[int]()
	for len(s) > 0 {
		x := s.Pop()
		if x.op {
			switch x.value {
			case 0:
				sum := 0
				for _, e := range valStack {
					sum += e
				}
				s.Push(item{sum, false})
				valStack.Reset()
			case 1:
				product := 1
				for _, e := range valStack {
					product *= e
				}
				s.Push(item{product, false})
				valStack.Reset()
			case 2:
				min := math.MaxInt
				for _, e := range valStack {
					if e < min {
						min = e
					}
				}
				s.Push(item{min, false})
				valStack.Reset()
			case 3:
				max := 0
				for _, e := range valStack {
					if e > max {
						max = e
					}
				}
				s.Push(item{max, false})
				valStack.Reset()
			case 5:
				first, second := valStack.Pop(), valStack.Pop()
				greater := 0
				if first > second {
					greater = 1
				}
				s.Push(item{greater, false})
				valStack.Reset()
			case 6:
				first, second := valStack.Pop(), valStack.Pop()
				less := 0
				if first < second {
					less = 1
				}
				s.Push(item{less, false})
				valStack.Reset()
			case 7:
				first, second := valStack.Pop(), valStack.Pop()
				equal := 0
				if first == second {
					equal = 1
				}
				s.Push(item{equal, false})
				valStack.Reset()
			}
		} else {
			valStack.Push(x.value)
		}
//		fmt.Println(s, valStack)
	}
	return valStack[0]
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
