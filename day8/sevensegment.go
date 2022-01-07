package day8

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func FindDigits(entries []string) int {
	res := 0
	for _, entry := range entries {
		pattern, output := parseEntry(entry)
		pattern = append(pattern, output...)
		digitMap := findDigits(pattern)
		digit := decodeDigits(digitMap, output)
		res += digit
	}
	return res
}

func findDigits(s []string) map[int]string {
	d := make(map[int]string)
	findOneFourSevenEight(d, s)
	findThree(d, s)
	findNine(d, s)
	findFive(d, s)
	findTwo(d, s)
	findZero(d, s)
	findSix(d, s)
	//	tryPrintDigits(d, s)
	return d
}

func findOneFourSevenEight(d map[int]string, pattern []string) {
	for _, p := range pattern {
		switch len(p) {
		case 2:
			d[1] = p
		case 3:
			d[7] = p
		case 4:
			d[4] = p
		case 7:
			d[8] = p
		}
	}
}

func findThree(d map[int]string, pattern []string) {
	seven, ok := d[7]
	if !ok {
		log.Fatal("error: 7 is not found, and 7 needed to find 3")
	}
	for _, p := range pattern {
		if len(p) != 5 {
			continue
		}
		if containsEach(p, seven) {
			d[3] = p
			break
		}
	}
}

// v2
func findFive(d map[int]string, pattern []string) {
	nine, ok := d[9]
	ninePerms := nCr(nine, 6, 5)
	if !ok {
		log.Fatal("error: 9 not found, and 9 needed to find 5")
	}
loop:
	for _, p := range pattern {
		if len(p) != 5 {
			continue
		}
		for _, c := range ninePerms {
			if c == p && d[3] != p {
				d[5] = p
				break loop
			}
		}
	}
}

func findTwo(d map[int]string, pattern []string) {
	three, ok := d[3]
	if !ok {
		log.Fatal("error: 3 not found, and 3 needed to find 2")
	}
	five, ok := d[5]
	if !ok {
		log.Fatal("error: 5 not found, and 5 needed to find 2")
	}

	for _, p := range pattern {
		if len(p) != 5 {
			continue
		}

		if !(p == three) && !(p == five) {
			d[2] = p
			break
		}

	}
}

func findNine(d map[int]string, pattern []string) {
	three, ok := d[3]
	if !ok {
		log.Fatal("error: 3 not found, and 3 needed to find 9")
	}
	for _, p := range pattern {
		if len(p) != 6 {
			continue
		}
		if containsEach(p, three) {
			d[9] = p
			break
		}

	}
}

func findZero(d map[int]string, pattern []string) {
	seven, ok := d[7]
	if !ok {
		log.Fatal("error: 7 not found, and 7 needed to find 0")
	}
	nine, ok := d[9]
	if !ok {
		log.Fatal("error: 9 not found, and 9 needed to find 0")
	}
	for _, p := range pattern {
		if len(p) != 6 || p == nine {
			continue
		}
		if containsEach(p, seven) {
			d[0] = p
			break
		}
	}
}

func findSix(d map[int]string, pattern []string) {
	nine, ok := d[9]
	if !ok {
		log.Fatal("error: 9 not found, and 9 needed to find 6")
	}
	zero, ok := d[0]
	if !ok {
		log.Fatal("error: 0 not found, and 0 needed to find 6")
	}

	for _, p := range pattern {
		if len(p) != 6 {
			continue
		}

		if !(p == nine) && !(p == zero) {
			d[6] = p
			break
		}

	}
}

func tryPrintDigits(d map[int]string, pattern []string) {
	r := reverseMap(d)
	for _, p := range pattern {
		digit, ok := r[p]
		if ok {
			fmt.Printf(" %v  ", digit)
		} else {
			fmt.Printf(" %v ", p)
		}
	}
	fmt.Print("\n")
}

func decodeDigits(d map[int]string, output []string) int {
	res := float64(0)
	r := reverseMap(d)
	l := len(output)
	for i := 0; i < l; i++ {
		d := r[output[l-1-i]]
		res += float64(d) * math.Pow(10, float64(i))
	}
	return int(res)
}

func reverseMap(d map[int]string) map[string]int {
	res := make(map[string]int)
	for k, v := range d {
		res[v] = k
	}
	return res
}

// containsEach checks whether str cotains each character in s
// (not necessarily sequentially)
func containsEach(str, s string) bool {
	for _, letter := range s {
		if !strings.Contains(str, string(letter)) {
			return false
		}
	}
	return true
}

func nCr(s string, n int, r int) []string {
	res := make([]string, 0)
	if len(s) != n {
		log.Fatal("s length does not match n")
	}
	if n == r {
		res = append(res, s)
		return res
	}
	if n-r == 1 {
		//first char
		res = append(res, s[1:])

		//middle chars
		for i := 0; i < n-2; i++ {
			curr := s[:i+1] + s[i+2:]
			res = append(res, curr)
		}
		//last char
		res = append(res, s[:len(s)-1])

		return res
	}
	return nil
}

func parseEntry(s string) ([]string, []string) {
	digits := strings.Split(s, " ")
	pattern, output := make([]string, 0), make([]string, 0)
	pipe := false

	for _, d := range digits {
		if d == "|" {
			pipe = true
			continue
		}
		if !pipe {
			pattern = append(pattern, sortString(d))
		} else {
			output = append(output, sortString(d))
		}
	}
	return pattern, output
}
