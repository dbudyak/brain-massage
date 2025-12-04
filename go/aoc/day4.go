package main

import (
	"strings"
)

func SolveDay4Part1(input string) int {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	totalAccessible := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {

			if grid[row][column] != '@' {
				continue
			}

			nw, _ := safeGet(grid, row-1, column-1)
			n, _ := safeGet(grid, row-1, column)
			ne, _ := safeGet(grid, row-1, column+1)
			e, _ := safeGet(grid, row, column+1)
			se, _ := safeGet(grid, row+1, column+1)
			s, _ := safeGet(grid, row+1, column)
			sw, _ := safeGet(grid, row+1, column-1)
			w, _ := safeGet(grid, row, column-1)

			rollCount := 0
			if nw == '@' {
				rollCount++
			}
			if n == '@' {
				rollCount++
			}
			if ne == '@' {
				rollCount++
			}
			if e == '@' {
				rollCount++
			}
			if se == '@' {
				rollCount++
			}
			if s == '@' {
				rollCount++
			}
			if sw == '@' {
				rollCount++
			}
			if w == '@' {
				rollCount++
			}

			if rollCount < 4 {
				totalAccessible++
			}
		}
	}

	return totalAccessible
}

func SolveDay4Part2(input string) int {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	totalAccessible := 0

	for i := 0; i < 100; i++ {
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid[row]); column++ {

				if grid[row][column] != '@' {
					continue
				}

				nw, _ := safeGet(grid, row-1, column-1)
				n, _ := safeGet(grid, row-1, column)
				ne, _ := safeGet(grid, row-1, column+1)
				e, _ := safeGet(grid, row, column+1)
				se, _ := safeGet(grid, row+1, column+1)
				s, _ := safeGet(grid, row+1, column)
				sw, _ := safeGet(grid, row+1, column-1)
				w, _ := safeGet(grid, row, column-1)

				rollCount := 0
				if nw == '@' {
					rollCount++
				}
				if n == '@' {
					rollCount++
				}
				if ne == '@' {
					rollCount++
				}
				if e == '@' {
					rollCount++
				}
				if se == '@' {
					rollCount++
				}
				if s == '@' {
					rollCount++
				}
				if sw == '@' {
					rollCount++
				}
				if w == '@' {
					rollCount++
				}

				if rollCount < 4 {
					totalAccessible++
					grid[row][column] = 'x'
				}
			}
		}
	}

	return totalAccessible
}

func safeGet(grid [][]rune, i, j int) (rune, bool) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) {
		return grid[i][j], true
	}
	return 0, false
}
