package main

import (
	"fmt"
	"strings"

	"aoc.test/io"
)

const TODAY = 9

func main() {
	days := map[int]string{
		1: "Day 1: Secret Entrance",
		2: "Day 2: Gift Shop",
		3: "Day 3: Lobby",
		4: "Day 4: Printing Department",
		5: "Day 5: Printing Department",
		6: "Day 6: Trash Compactor",
		7: "Day 7: Laboratories",
		8: "Day 8: Playground",
		9: "Day 9: Movie Theater",
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
	case 4:
		{
			part1 := SolveDay4Part1(input)
			fmt.Printf("%d rolls of paper\n", part1)

			part2 := SolveDay4Part2(input)
			fmt.Printf("%d rolls of paper\n", part2)
		}
	case 5:
		{
			part1 := SolveDay5Part1(input)
			fmt.Printf("\n%d fresh ids", part1)

			part2 := SolveDay5Part2(input)
			fmt.Printf("\n%d fresh ids", part2)
		}
	case 6:
		{
			part1 := SolveDay6Part1(input)
			fmt.Printf("\nGrand total p1 is %d", part1)

			part2 := SolveDay6Part2(io.MustReadInput(TODAY))
			fmt.Printf("\nGrand total p2 is %d", part2)
		}
	case 7:
		{
			part1 := SolveDay7Part1(input)
			fmt.Printf("\nBeam split %d times", part1)

			part2 := SolveDay7Part2(input)
			fmt.Printf("\n%d timelines", part2)
		}
	case 8:
		{
			part1 := SolveDay8Part1(input)
			fmt.Printf("\nCircuit size is %d", part1)

			part2 := SolveDay8Part2(input)
			fmt.Printf("\nCircuit distance is %d", part2)
		}
	case 9:
		{
			part1 := SolveDay9Part1(input)
			fmt.Printf("\nLargest area is %d", part1)
		}
	}

}
