package main

import (
	"strings"
)

type Position struct {
	row, col int
}

func SolveDay7Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]rune, len(lines))
	startCol := -1
	for i, line := range lines {
		grid[i] = []rune(line)
		for j, ch := range line {
			if ch == 'S' {
				startCol = j
			}
		}
	}

	hitSplitters := make(map[Position]bool)
	beams := []Position{{row: 0, col: startCol}}
	seen := make(map[Position]bool)

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		if seen[beam] {
			continue
		}
		seen[beam] = true

		for row := beam.row; row < len(grid); row++ {
			col := beam.col
			if col < 0 || col >= len(grid[row]) {
				break
			}

			cell := grid[row][col]
			if cell == '^' {
				pos := Position{row: row, col: col}
				hitSplitters[pos] = true

				if row+1 < len(grid) {
					if col-1 >= 0 {
						beams = append(beams, Position{row: row + 1, col: col - 1})
					}
					if col+1 < len(grid[row+1]) {
						beams = append(beams, Position{row: row + 1, col: col + 1})
					}
				}
				break
			}
		}
	}

	return len(hitSplitters)
}
