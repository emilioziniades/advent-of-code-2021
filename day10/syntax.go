package day10

var open = map[rune]bool{
	'(': true,
	'{': true,
	'<': true,
	'[': true,
}
var clos = map[rune]bool{
	')': true,
	'}': true,
	'>': true,
	']': true,
}
var flip = map[rune]rune{
	'(': ')',
	')': '(',
	'[': ']',
	']': '[',
	'{': '}',
	'}': '{',
	'<': '>',
	'>': '<',
}
var score = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func ErrorScore(fs []string) int {
	eScore := 0
loop:
	for _, f := range fs {
		//		fmt.Println(i)
		stack := make([]rune, 0)
		for _, s := range f {
			if open[s] {
				stack = append(stack, s)
			}

			if clos[s] {
				if stack[len(stack)-1] != flip[s] {
					//					fmt.Printf("%q : %q ", stack, s)
					//					fmt.Printf(" Expected %q, but found %q instead\n", flip[stack[len(stack)-1]], s)
					eScore += score[s]
					continue loop
				} else {
					// has matching closing brace
					//					fmt.Printf("%q : %q\n", stack, s)
					stack = stack[:len(stack)-1]
				}

			}
		}
	}
	return eScore
}
