package main

import (
	"strconv"
	"strings"
)

func SolveDay6Part1(input string) int {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, 0)
	for _, line := range lines {
		lineArr := strings.Fields(line)
		matrix = append(matrix, lineArr)
	}

	grandTotal := 0

	ops := matrix[len(lines)-1]

	for opId := 0; opId < len(ops); opId++ {
		switch ops[opId] {
		case "*":
			{
				localTotal := 1
				for rowId := 0; rowId < len(lines)-1; rowId++ {
					num, err := strconv.Atoi(matrix[rowId][opId])
					if err != nil {
						panic(err)
					}
					localTotal *= num
				}
				grandTotal += localTotal
			}
		case "+":
			{
				localTotal := 0
				for rowId := 0; rowId < len(lines)-1; rowId++ {
					num, err := strconv.Atoi(matrix[rowId][opId])
					if err != nil {
						panic(err)
					}
					localTotal += num
				}
				grandTotal += localTotal
			}
		}

	}
	return grandTotal
}

func SolveDay6Part2(input string) int64 {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return 0
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	opsLine := grid[len(grid)-1]

	var grandTotal int64 = 0
	currentOp := ""
	numbers := make([]int64, 0)

	for colId := 0; colId < len(opsLine); colId++ {
		ch := opsLine[colId]

		if ch == '*' || ch == '+' {
			currentOp = string(ch)
			numbers = make([]int64, 0)
		}

		digitStr := ""
		for rowId := 0; rowId < len(grid)-1; rowId++ {
			if colId < len(grid[rowId]) && grid[rowId][colId] >= '0' && grid[rowId][colId] <= '9' {
				digitStr += string(grid[rowId][colId])
			}
		}
		if digitStr != "" {
			val, err := strconv.ParseInt(digitStr, 10, 64)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, val)

			if colId == len(opsLine)-1 {
				grandTotal += applyOperation(numbers, currentOp)
			}
		} else {
			grandTotal += applyOperation(numbers, currentOp)
		}
	}

	return grandTotal
}

func applyOperation(numbers []int64, op string) int64 {
	if len(numbers) == 0 {
		return 0
	}

	var result int64
	switch op {
	case "*":
		result = 1
		for _, num := range numbers {
			result *= num
		}
	case "+":
		result = 0
		for _, num := range numbers {
			result += num
		}
	}
	return result
}
