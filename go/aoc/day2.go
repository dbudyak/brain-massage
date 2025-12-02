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
	for i := 0; i < len(ranges)-1; i++ {
		minMax := strings.Split(ranges[i], "-")
		start, _ := strconv.Atoi(minMax[0])
		end, _ := strconv.Atoi(minMax[1])

		for id := start; id <= end; id++ {
			if isInvalid(id) {
				fmt.Printf("id=%d\n", id)
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
	return false
}
