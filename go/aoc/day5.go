package main

import (
	"sort"
	"strconv"
	"strings"
)

func SolveDay5Part1(input string) int {
	lines := strings.Split(input, "\n\n")

	rangesStr := strings.Split(lines[0], "\n")
	idsStr := strings.Split(lines[1], "\n")

	ranges := make([][]int, 0)
	for _, freshRange := range rangesStr {
		minMax := strings.Split(freshRange, "-")
		start, err := strconv.Atoi(minMax[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(minMax[1])
		if err != nil {
			panic(err)
		}

		ranges = append(ranges, []int{start, end})
	}

	freshIds := make(map[int]struct{}, 0)

	for _, availableId := range idsStr {
		freshId, err := strconv.Atoi(availableId)
		if err != nil {
			panic(err)
		}

		for _, freshRange := range ranges {
			_, exists := freshIds[freshId]
			if freshId >= freshRange[0] && freshId <= freshRange[1] && !exists {
				freshIds[freshId] = struct{}{}
			}
		}
	}

	return len(freshIds)
}

func SolveDay5Part2(input string) int {
	lines := strings.Split(input, "\n\n")
	rangesStr := strings.Split(lines[0], "\n")

	ranges := make([][2]int, 0)
	for _, freshRange := range rangesStr {
		freshRange = strings.TrimSpace(freshRange)
		if freshRange == "" {
			continue
		}

		minMax := strings.Split(freshRange, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(minMax[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(minMax[1]))
		ranges = append(ranges, [2]int{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := make([][2]int, 0)
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}

		last := &merged[len(merged)-1]

		if r[0] <= last[1]+1 {
			if r[1] > last[1] {
				last[1] = r[1]
			}
		} else {
			merged = append(merged, r)
		}
	}

	total := 0
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}

	return total
}
