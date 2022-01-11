package day10

import (
	"sort"
)

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

var scoreC = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func ErrorScore(fs []string) int {
	eScore := 0
loop:
	for _, f := range fs {
		stack := make([]rune, 0)
		for _, s := range f {
			if open[s] {
				stack = append(stack, s)
			}

			if clos[s] {
				if stack[len(stack)-1] != flip[s] {
					// incorrect closing brace
					eScore += score[s]
					continue loop
				} else {
					// has matching closing brace
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return eScore
}

func CompletionScore(fs []string) int {
	cScores := make([]int, 0)
loop:
	for _, f := range fs {
		stack := make([]rune, 0)
		for _, s := range f {
			if open[s] {
				stack = append(stack, s)
			}

			if clos[s] {
				if stack[len(stack)-1] != flip[s] {
					// no matching closing brace
					continue loop
				} else {
					// has matching closing brace
					stack = stack[:len(stack)-1]
				}
			}
		}
		reverse(stack)
		cScores = append(cScores, countScore(stack))
	}
	sort.Ints(cScores)
	return cScores[((len(cScores)+1)/2)-1]
}

func reverse(rs []rune) {
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
}

func countScore(rs []rune) int {
	score := 0
	for _, r := range rs {
		score *= 5
		score += scoreC[r]
	}
	return score
}
