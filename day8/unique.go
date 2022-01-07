package day8

func UniqueDigits(entries []string) (n int) {
	for _, entry := range entries {
		_, output := parseEntry(entry)
		c := countUnique(output)
		n += c
	}
	return
}

func countUnique(pattern []string) int {
	n := 0
	for _, p := range pattern {
		if l := len(p); ((l >= 2) && (l <= 4)) || l == 7 {
			//			fmt.Println(p)
			n += 1
		}
	}
	return n
}
