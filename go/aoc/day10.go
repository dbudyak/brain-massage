package main

import (
	"fmt"
	"regexp"
	"strings"
)

func SolveDay10Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	totalPresses := 0

	for _, line := range lines {
		presses := solveMachine(line)
		totalPresses += presses
	}

	return totalPresses
}

func solveMachine(line string) int {
	targetRegex := regexp.MustCompile(`\[([.#]+)\]`)
	targetMatch := targetRegex.FindStringSubmatch(line)
	target := targetMatch[1]
	numLights := len(target)

	buttonRegex := regexp.MustCompile(`\(([0-9,]+)\)`)
	buttonMatches := buttonRegex.FindAllStringSubmatch(line, -1)

	buttons := make([][]int, len(buttonMatches))
	for i, match := range buttonMatches {
		numStrs := strings.Split(match[1], ",")
		buttons[i] = make([]int, 0)
		for _, numStr := range numStrs {
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			buttons[i] = append(buttons[i], num)
		}
	}

	fmt.Printf("\nTarget: %s (%d lights)\n", target, numLights)
	fmt.Printf("Buttons: %v\n", buttons)

	matrix := make([][]int, numLights)
	for i := range matrix {
		matrix[i] = make([]int, len(buttons)+1)
	}

	for buttonIdx, button := range buttons {
		for _, lightIdx := range button {
			if lightIdx < numLights {
				matrix[lightIdx][buttonIdx] = 1
			}
		}
	}

	for i, ch := range target {
		if ch == '#' {
			matrix[i][len(buttons)] = 1
		}
	}

	fmt.Printf("Matrix before solving:\n")
	for i, row := range matrix {
		fmt.Printf("Light %d: %v\n", i, row)
	}

	solution := gaussianEliminationGF2(matrix, len(buttons))

	fmt.Printf("Solution: %v\n", solution)

	if solution == nil {
		return -1
	}

	count := 0
	for i, val := range solution {
		if val == 1 {
			fmt.Printf("Press button %d: %v\n", i, buttons[i])
			count++
		}
	}

	fmt.Printf("Total presses: %d\n", count)
	return count
}

func gaussianEliminationGF2(matrix [][]int, numVars int) []int {
	rows := len(matrix)
	cols := numVars

	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols+1)
		copy(m[i], matrix[i])
	}

	pivotCols := make([]int, 0)

	pivotRow := 0
	for col := 0; col < cols && pivotRow < rows; col++ {
		found := false
		for row := pivotRow; row < rows; row++ {
			if m[row][col] == 1 {
				m[pivotRow], m[row] = m[row], m[pivotRow]
				found = true
				break
			}
		}

		if !found {
			continue
		}

		pivotCols = append(pivotCols, col)

		for row := 0; row < rows; row++ {
			if row != pivotRow && m[row][col] == 1 {
				for c := 0; c <= cols; c++ {
					m[row][c] ^= m[pivotRow][c]
				}
			}
		}

		pivotRow++
	}

	for row := pivotRow; row < rows; row++ {
		allZero := true
		for col := 0; col < cols; col++ {
			if m[row][col] == 1 {
				allZero = false
				break
			}
		}
		if allZero && m[row][cols] == 1 {
			return nil
		}
	}

	freeVars := make([]int, 0)
	pivotSet := make(map[int]bool)
	for _, col := range pivotCols {
		pivotSet[col] = true
	}
	for col := 0; col < cols; col++ {
		if !pivotSet[col] {
			freeVars = append(freeVars, col)
		}
	}

	numFree := len(freeVars)
	bestSolution := []int(nil)
	minCount := cols + 1

	for mask := 0; mask < (1 << numFree); mask++ {
		solution := make([]int, cols)

		for i, freeVar := range freeVars {
			if (mask & (1 << i)) != 0 {
				solution[freeVar] = 1
			}
		}

		for i := len(pivotCols) - 1; i >= 0; i-- {
			col := pivotCols[i]
			val := m[i][cols]
			for j := col + 1; j < cols; j++ {
				val ^= (m[i][j] * solution[j])
			}
			solution[col] = val
		}

		count := 0
		for _, v := range solution {
			count += v
		}

		if count < minCount {
			minCount = count
			bestSolution = make([]int, cols)
			copy(bestSolution, solution)
		}
	}

	return bestSolution
}
