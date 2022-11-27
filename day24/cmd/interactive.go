package main

import (
	"fmt"
	"log"

	"github.com/emilioziniades/adventofcode2021/day24"
)

func main() {
	stepInteractive()
}

func stepInteractive() {
	changingValues := day24.GetChangingValues("input.txt")
	var z int
	for n := 0; n < 14; n++ {
		vals := changingValues[n]
		fmt.Printf("z%d: %d\n", n, z)
		printOptions(z, vals.A, vals.B, vals.C)
		i := scanDigit(fmt.Sprintf("i%d: ", n))
		z = day24.Step(z, i, vals.A, vals.B, vals.C)
	}
	fmt.Printf("z: %d\n", z)
}

func printOptions(z, a, b, c int) {
	for i := 1; i <= 9; i++ {
		zn := day24.Step(z, i, a, b, c)
		fmt.Printf("i = %d\tz = %d\n", i, zn)
	}
}

func scanDigit(msg string) int {
	fmt.Print(msg)
	var n int
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		log.Fatal(err)
	}
	if n > 9 || n < 1 {
		log.Fatal("Must have 1 <= n <= 9")
	}

	return n
}
