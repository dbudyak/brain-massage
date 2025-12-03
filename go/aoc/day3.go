package main

import (
	"strings"
)

func SolveDay3Part1(input string) int {
	banks := strings.Split(input, "\n")
	total := 0
	for _, bank := range banks {
		total += GetMaximumJoltage(bank)
	}

	return total
}

func GetMaximumJoltage(bank string) int {
	maxJoltage := 0

	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			digit1 := int(bank[i] - '0')
			digit2 := int(bank[j] - '0')
			joltage := digit1*10 + digit2

			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
	}
	return maxJoltage
}

func SolveDay3Part2(input string) int {
	banks := strings.Split(input, "\n")
	total := 0
	for _, bank := range banks {
		total += GetMaximumJoltagePart2(bank, 12)
	}

	return total
}

func GetMaximumJoltagePart2(bank string, numBatteries int) int {
	if len(bank) < numBatteries {
		return 0
	}

	result := 0
	lastIndex := -1

	for picked := 0; picked < numBatteries; picked++ {
		remainingNeeded := numBatteries - picked - 1
		maxPos := len(bank) - remainingNeeded - 1
		maxDigit := -1
		maxDigitPos := -1
		for i := lastIndex + 1; i <= maxPos; i++ {
			digit := int(bank[i] - '0')
			if digit > maxDigit {
				maxDigit = digit
				maxDigitPos = i
			}
		}

		result = result*10 + maxDigit
		lastIndex = maxDigitPos
	}

	return result
}
