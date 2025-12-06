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

/*
func SolveDay6Part2(input string) int {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, 0)
	for _, line := range lines {
		lineArr := strings.Fields(line)
		matrix = append(matrix, lineArr)
	}

	grandTotal := 0

	ops := matrix[len(lines)-1]

	for opId := 0; opId < len(ops); opId++ {
		numbersStr := make([]string, 0)
		for rowId := 0; rowId < len(lines)-1; rowId++ {
			numbersStr = append(numbersStr, matrix[rowId][opId])
		}
		grandTotal += DoSpecialMath(numbersStr, ops[opId])
	}
	return grandTotal
}

func getRunesByNumberId(numId int, numbers []string) []rune {
	if len(numbers) == 0 {
		return []rune{}
	}

	if numId < len(numbers) {
		return []rune(numbers[numId])
	} else {
		return []rune{}
	}
}

func DoSpecialMath(numbers []string, op string) int {

	num1 := getRunesByNumberId(0, numbers)
	num2 := getRunesByNumberId(1, numbers)
	num3 := getRunesByNumberId(2, numbers)
	num4 := getRunesByNumberId(3, numbers)

	result1 := getRuneById(3, num1) + getRuneById(3, num2) + getRuneById(3, num3) + getRuneById(3, num4)
	result2 := getRuneById(2, num1) + getRuneById(2, num2) + getRuneById(2, num3) + getRuneById(2, num4)
	result3 := getRuneById(1, num1) + getRuneById(1, num2) + getRuneById(1, num3) + getRuneById(1, num4)
	result4 := getRuneById(0, num1) + getRuneById(0, num2) + getRuneById(0, num3) + getRuneById(0, num4)

	println(result1)
	println(result2)
	println(result3)
	println(result4)

	switch op {
	case "*":
		{

			res1 := getIntOrElse(result1, 1)
			res2 := getIntOrElse(result2, 1)
			res3 := getIntOrElse(result3, 1)
			res4 := getIntOrElse(result4, 1)

			return res1 * res2 * res3 * res4
		}
	case "+":
		{
			res1 := getIntOrElse(result1, 0)
			res2 := getIntOrElse(result2, 0)
			res3 := getIntOrElse(result3, 0)
			res4 := getIntOrElse(result4, 0)

			return res1 + res2 + res3 + res4
		}
	default:
		return 0
	}

}

func getIntOrElse(numStr string, defaultVal int) int {
	if numStr != "" {
		res, err := strconv.Atoi(numStr)
		if err == nil {
			return res
		}
	}
	return defaultVal
}

func getRuneById(id int, runes []rune) string {
	if id < len(runes) {
		return string(runes[id])
	} else {
		return ""
	}
}

*/
