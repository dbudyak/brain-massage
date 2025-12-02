package main

import (
	"fmt"
	"strings"

	"aoc.test/io"
)

func main() {

	days := map[int]string{
		1: "Day 1: Secret Entrance",
		2: "Day 2: Gift Shop",
	}
	today := 2

	fmt.Println(days[today])
	fmt.Println()

	input := io.MustReadInput(today)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Printf("Read %d lines from input1.txt\n", len(lines))
	fmt.Println()

	switch today {
	case 1:
		{
			part1 := SolveDay1(input)
			fmt.Printf("Password (Part 1): %d\n", part1)
			fmt.Println()

			part2 := SolveDay1Part2(input)
			fmt.Printf("Password (Part 2): %d\n", part2)
		}
	case 2:
		{
			part1 := SolveDay2Part1(input)
			fmt.Printf("Sum of invalid IDs: %d\n", part1)
		}

	}
}
