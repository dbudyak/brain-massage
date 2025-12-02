package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolveDay2(input string, isInvalid func(id int) bool) int {
	totalSum := 0

	ranges := strings.Split(input, ",")

	for i := 0; i < len(ranges); i++ {
		minMax := strings.Split(ranges[i], "-")
		start, err1 := strconv.Atoi(strings.TrimSpace(minMax[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(minMax[1]))

		if err1 != nil || err2 != nil {
			fmt.Printf("Error parsing range: %v-%v, errors: %v, %v\n", minMax[0], minMax[1], err1, err2)
			continue
		}

		for id := start; id <= end; id++ {
			if isInvalid(id) {
				totalSum += id
			}
		}

	}
	return totalSum
}

func IsInvalidPart1(id int) bool {
	digits := len(fmt.Sprintf("%d", id))
	if digits%2 != 0 {
		return false
	}

	divider := int(math.Pow10(digits / 2))
	left := id / divider
	right := id % divider

	return left == right
}

func IsInvalidPart2(id int) bool {
	if id < 10 {
		return false
	}

	numberStr := strconv.Itoa(id)
	maxSubstringLength := len(numberStr) / 2

	for i := 0; i < maxSubstringLength; i++ {

		substring := numberStr[0 : i+1]
		result := strings.ReplaceAll(numberStr, substring, "")

		if result == "" {
			return true
		}
	}
	return false
}
