package main

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

type Rectangle struct {
	p1, p2 Point
}

func SolveDay9Part1(input string) int {
	recsStr := strings.Split(input, "\n")

	points := make([]Point, len(recsStr))
	for i := 0; i < len(recsStr); i++ {
		result := strings.Split(recsStr[i], ",")
		x, err := strconv.ParseFloat(result[0], 64)
		if err != nil {
			panic(err)
		}
		y, err := strconv.ParseFloat(result[1], 64)
		if err != nil {
			panic(err)
		}
		points[i] = Point{x: x, y: y}
	}

	maxSize := 0.0

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			rectangle := Rectangle{p1: points[i], p2: points[j]}
			size := rectangle.size()
			if size > maxSize {
				maxSize = size
			}
		}
	}

	return int(maxSize)
}

func (rect Rectangle) size() float64 {
	width := math.Abs(rect.p2.x-rect.p1.x) + 1
	height := math.Abs(rect.p2.y-rect.p1.y) + 1
	return width * height
}
