package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.test/io"
)

const (
	left  int = -1
	right int = 1
)

func main() {
	fmt.Println("Day 1: Secret Entrance")
	fmt.Println()

	input := io.MustReadInput(1)

	lines := strings.Split(strings.TrimSpace(input), "\n")
	fmt.Printf("Read %d lines from input1.txt\n", len(lines))
	fmt.Println()

	part1 := SolveDay1(input)
	fmt.Printf("Password (Part 1): %d\n", part1)
	fmt.Println()

	part2 := SolveDay1Part2(input)
	fmt.Printf("Password (Part 2): %d\n", part2)
}

func Rotate(position int, distance int, direction int) int {
	const dialSize = 100
	newPosition := position + (direction * distance)
	return ((newPosition % dialSize) + dialSize) % dialSize
}

func CountZeroCrossings(position int, distance int, direction int) int {
	if direction == right {
		return (position + distance) / 100
	} else {
		if position == 0 {
			return distance / 100
		}
		if distance >= position {
			return 1 + (distance-position)/100
		}
		return 0
	}
}

func SolveDay1(input string) int {
	const startPosition = 50
	position := startPosition
	zeroCount := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		direction := left
		if line[0] == 'R' {
			direction = right
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		position = Rotate(position, distance, direction)

		if position == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func SolveDay1Part2(input string) int {
	const startPosition = 50
	position := startPosition
	totalCount := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		direction := left
		if line[0] == 'R' {
			direction = right
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}

		crossings := CountZeroCrossings(position, distance, direction)
		totalCount += crossings

		position = Rotate(position, distance, direction)
	}

	return totalCount
}
