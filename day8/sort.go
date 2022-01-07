package day8

import "sort"

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(runeArr(r))
	return string(r)

}

type runeArr []rune

func (r runeArr) Len() int {
	return len(r)
}

func (r runeArr) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeArr) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
