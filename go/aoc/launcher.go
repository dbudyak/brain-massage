package main

import (
	"fmt"
	"strings"

	"aoc.test/io"
)

const TODAY = 3

func main() {
	days := map[int]string{
		1: "Day 1: Secret Entrance",
		2: "Day 2: Gift Shop",
		3: "Day 3: Lobby",
	}

	fmt.Println(days[TODAY])
	fmt.Println()

	input := strings.TrimSpace(io.MustReadInput(TODAY))

	lines := strings.Split(input, "\n")
	fmt.Printf("Read %d lines from input1.txt\n", len(lines))
	fmt.Println()

	switch TODAY {
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
			part1 := SolveDay2(input, IsInvalidPart1)
			fmt.Printf("Sum of invalid IDs for Part 1: %d\n", part1)

			part2 := SolveDay2(input, IsInvalidPart2)
			fmt.Printf("Sum of invalid IDs for Part 2: %d\n", part2)
		}
	case 3:
		{
			part1 := SolveDay3Part1(input)
			fmt.Printf("Max joltage for Part 1: %d\n", part1)

			part2 := SolveDay3Part2(input)
			fmt.Printf("Max joltage for Part 2: %d\n", part2)
		}
	}
}
