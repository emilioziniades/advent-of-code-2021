package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Move(commands []string) int {
	var x, y int
	for _, command := range commands {
		current := strings.Split(command, " ")
		magnitude, err := strconv.Atoi(current[1])
		if err != nil {
			log.Fatalf("move: strconv: %s\n", err)
		}

		switch direction := current[0]; direction {
		case "forward":
			x += magnitude
		case "down":
			y += magnitude
		case "up":
			y -= magnitude
		}
	}
	fmt.Printf("x: %d, y: %d, product: %d\n", x, y, x*y)
	return x * y
}

func MoveAim(commands []string) int {
	var x, y, aim int
	for _, command := range commands {
		current := strings.Split(command, " ")
		magnitude, err := strconv.Atoi(current[1])
		if err != nil {
			log.Fatalf("move: strconv: %s\n", err)
		}

		switch direction := current[0]; direction {
		case "forward":
			x += magnitude
			y += (magnitude * aim)
		case "down":
			aim += magnitude
		case "up":
			aim -= magnitude
		}
	}
	fmt.Printf("x: %d, y: %d, product: %d\n", x, y, x*y)
	return x * y
}
