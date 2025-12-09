package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type JBox struct {
	x, y, z float64
}

type Circuit struct {
	jboxes []JBox
}

func SolveDay8Part1(input string) int {
	jboxes := initJBoxArray(input)
	distances := make(map[JBox]map[JBox]float64)
	circuits := make([]Circuit, 0)

	for _, jboxi := range jboxes {
		for _, jboxj := range jboxes {
			setDistance(distances, jboxi, jboxj)
		}
	}

	excludeDistances := make([]float64, 0)

	maxIter := 10
	if len(jboxes) >= 21 {
		maxIter = 1000
	}

	for i := 0; i < maxIter; i++ {
		jb1, jb2, dist := findMinDistance(distances, excludeDistances)
		excludeDistances = append(excludeDistances, dist)

		if !areBothInCircuit(jb1, jb2, circuits) {
			idx1, exists1 := isInCircuit(jb1, circuits)
			idx2, exists2 := isInCircuit(jb2, circuits)

			if exists1 && exists2 {
				circuits[idx1].jboxes = append(circuits[idx1].jboxes, circuits[idx2].jboxes...)
				circuits = append(circuits[:idx2], circuits[idx2+1:]...)
			} else if exists1 {
				circuits[idx1].jboxes = append(circuits[idx1].jboxes, jb2)
			} else if exists2 {
				circuits[idx2].jboxes = append(circuits[idx2].jboxes, jb1)
			} else {
				circuits = append(circuits, Circuit{jboxes: []JBox{jb1, jb2}})
			}
		}
	}

	for _, jbox := range jboxes {
		_, exists := isInCircuit(jbox, circuits)
		if !exists {
			circuits = append(circuits, Circuit{jboxes: []JBox{jbox}})
		}
	}

	return multiplyThreeLargest(circuits)
}

func SolveDay8Part2(input string) int {
	jboxes := initJBoxArray(input)
	distances := make(map[JBox]map[JBox]float64)
	circuits := make([]Circuit, 0)

	for _, jboxi := range jboxes {
		for _, jboxj := range jboxes {
			setDistance(distances, jboxi, jboxj)
		}
	}

	excludeDistances := make([]float64, 0)
	var lastJb1, lastJb2 JBox

	for {
		jb1, jb2, dist := findMinDistance(distances, excludeDistances)
		excludeDistances = append(excludeDistances, dist)

		if !areBothInCircuit(jb1, jb2, circuits) {
			idx1, exists1 := isInCircuit(jb1, circuits)
			idx2, exists2 := isInCircuit(jb2, circuits)

			if exists1 && exists2 {
				circuits[idx1].jboxes = append(circuits[idx1].jboxes, circuits[idx2].jboxes...)
				circuits = append(circuits[:idx2], circuits[idx2+1:]...)
			} else if exists1 {
				circuits[idx1].jboxes = append(circuits[idx1].jboxes, jb2)
			} else if exists2 {
				circuits[idx2].jboxes = append(circuits[idx2].jboxes, jb1)
			} else {
				circuits = append(circuits, Circuit{jboxes: []JBox{jb1, jb2}})
			}

			lastJb1, lastJb2 = jb1, jb2
		}

		tempCircuits := make([]Circuit, len(circuits))
		copy(tempCircuits, circuits)
		for _, jbox := range jboxes {
			_, exists := isInCircuit(jbox, tempCircuits)
			if !exists {
				tempCircuits = append(tempCircuits, Circuit{jboxes: []JBox{jbox}})
			}
		}

		if len(tempCircuits) == 1 {
			break
		}
	}

	return int(lastJb1.x * lastJb2.x)
}

func multiplyThreeLargest(circuits []Circuit) int {
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].jboxes) > len(circuits[j].jboxes)
	})

	return len(circuits[0].jboxes) * len(circuits[1].jboxes) * len(circuits[2].jboxes)
}

func areBothInCircuit(jbox1 JBox, jbox2 JBox, circuits []Circuit) bool {
	for _, circuit := range circuits {
		if containsBothInCircuit(jbox1, jbox2, circuit) {
			return true
		}
	}
	return false
}

func initJBoxArray(input string) []JBox {
	jboxesStr := strings.Split(input, "\n")
	jboxes := make([]JBox, len(jboxesStr))
	for i, jboxStr := range jboxesStr {
		jboxCoords := strings.Split(jboxStr, ",")

		x, err := strconv.ParseFloat(jboxCoords[0], 64)
		if err != nil {
			panic(err)
		}

		y, err := strconv.ParseFloat(jboxCoords[1], 64)
		if err != nil {
			panic(err)
		}

		z, err := strconv.ParseFloat(jboxCoords[2], 64)
		if err != nil {
			panic(err)
		}

		jboxes[i] = JBox{x, y, z}
	}
	return jboxes
}

func setDistance(m map[JBox]map[JBox]float64, jbox1, jbox2 JBox) {
	dist := distance(jbox1, jbox2)

	if m[jbox1] == nil {
		m[jbox1] = make(map[JBox]float64)
	}
	if m[jbox2] == nil {
		m[jbox2] = make(map[JBox]float64)
	}
	m[jbox1][jbox2] = dist
	m[jbox2][jbox1] = dist
}

func findMinDistance(distances map[JBox]map[JBox]float64, skip []float64) (JBox, JBox, float64) {
	minDist := math.MaxFloat64
	var p1, p2 JBox

	for from, dists := range distances {
		for to, dist := range dists {
			if !from.equals(to) && dist < minDist && !contains(skip, dist) {
				minDist = dist
				p1, p2 = from, to
			}
		}
	}

	return p1, p2, minDist
}

func distance(j1, j2 JBox) float64 {
	dx := j2.x - j1.x
	dy := j2.y - j1.y
	dz := j2.z - j1.z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (jbox JBox) toString() string {
	return fmt.Sprintf("(%v, %v, %v)", jbox.x, jbox.y, jbox.z)
}

func (jbox JBox) equals(another JBox) bool {
	if jbox.x == another.x && jbox.y == another.y && jbox.z == another.z {
		return true
	}
	return false
}

func contains(distances []float64, distance float64) bool {
	for _, v := range distances {
		if v == distance {
			return true
		}
	}
	return false
}

func containsBothInCircuit(jbox1 JBox, jbox2 JBox, circuit Circuit) bool {
	match1, match2 := false, false
	for _, v := range circuit.jboxes {
		if v.equals(jbox1) {
			match1 = true
		}
		if v.equals(jbox2) {
			match2 = true
		}
	}
	return match1 && match2
}

func containsInCircuit(jbox JBox, circuit Circuit) bool {
	for _, v := range circuit.jboxes {
		if v.equals(jbox) {
			return true
		}
	}
	return false
}

func isInCircuit(jbox JBox, circuits []Circuit) (int, bool) {
	for i, circuit := range circuits {
		if containsInCircuit(jbox, circuit) {
			return i, true
		}
	}
	return -1, false
}
