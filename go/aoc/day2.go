package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolveDay2Part1(input string) int {

	totalSum := 0

	ranges := strings.Split(input, ",")
	for i := 0; i < len(ranges)-1; i++ {
		minMax := strings.Split(ranges[i], "-")
		start, _ := strconv.Atoi(minMax[0])
		end, _ := strconv.Atoi(minMax[1])

		for id := start; id <= end; id++ {
			if IsInvalid(id) {
				fmt.Printf("id=%d\n", id)
				totalSum += id
			}
		}

	}
	return totalSum
}

func Sum(integers [][]int) int {
	total := 0
	for _, i := range integers {
		for _, j := range i {
			total += j
		}
	}
	return total
}

func IsInvalid(id int) bool {
	digits := len(fmt.Sprintf("%d", id))
	if digits%2 != 0 {
		return false
	}

	divider := int(math.Pow10(digits / 2))
	left := id / divider
	right := id % divider

	return left == right
}

func GetInvalidIds(start int, end int) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= end; i++ {
		if IsInvalid(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
