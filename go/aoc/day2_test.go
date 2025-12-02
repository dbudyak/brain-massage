package main

import (
	"reflect"
	"testing"
)

func TestSolveDay2_IsInvalidForPart1(t *testing.T) {
	valid := 12312

	if IsInvalidPart1(valid) {
		t.Errorf("%v is invalid", valid)
	}

	valid = 1

	if IsInvalidPart1(valid) {
		t.Errorf("%v is invalid", valid)
	}

	invalid := 1212
	if !IsInvalidPart1(invalid) {
		t.Errorf("%v is valid", valid)
	}
}

func TestSolveDay2_InvalidIdsPart1(t *testing.T) {

	tests := []struct {
		start      int
		end        int
		invalidIds []int
	}{
		{start: 11, end: 22, invalidIds: []int{11, 22}},
		{start: 95, end: 115, invalidIds: []int{99}},
		{start: 998, end: 1012, invalidIds: []int{1010}},
		{start: 1188511880, end: 1188511890, invalidIds: []int{1188511885}},
		{start: 1698522, end: 1698528, invalidIds: []int{}},
		{start: 446443, end: 446449, invalidIds: []int{446446}},
		{start: 38593856, end: 38593862, invalidIds: []int{38593859}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(getInvalidIds(test.start, test.end, IsInvalidPart1), test.invalidIds) {
			t.Errorf("GetInvalidIds(%v, %v) is invalid", test.start, test.end)
		}
	}
}

func TestSolveDay2_SolveDay2Part1(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	got := SolveDay2(input, IsInvalidPart1)
	want := 1227775554

	if got != want {
		t.Errorf("No stars today. SolveDay2 Part 2(%v) = %v, want %v", input, got, want)
	}
}

func TestSolveDay2_IsInvalidForPart2(t *testing.T) {
	valid := 12312

	if IsInvalidPart2(valid) {
		t.Errorf("%v is invalid", valid)
	}

	valid = 1

	if IsInvalidPart2(valid) {
		t.Errorf("%v is invalid", valid)
	}

	invalid := 1212
	if !IsInvalidPart2(invalid) {
		t.Errorf("%v is valid", valid)
	}

	invalid = 123123123
	if !IsInvalidPart2(invalid) {
		t.Errorf("%v is valid", valid)
	}
}

func TestSolveDay2_InvalidIdsPart2(t *testing.T) {

	tests := []struct {
		start      int
		end        int
		invalidIds []int
	}{
		{start: 11, end: 22, invalidIds: []int{11, 22}},
		{start: 95, end: 115, invalidIds: []int{99, 111}},
		{start: 998, end: 1012, invalidIds: []int{999, 1010}},
		{start: 1188511880, end: 1188511890, invalidIds: []int{1188511885}},
		{start: 222220, end: 222224, invalidIds: []int{222222}},
		{start: 1698522, end: 1698528, invalidIds: []int{}},
		{start: 38593856, end: 38593862, invalidIds: []int{38593859}},
		{start: 565653, end: 565659, invalidIds: []int{565656}},
		{start: 824824821, end: 824824827, invalidIds: []int{824824824}},
		{start: 2121212118, end: 2121212124, invalidIds: []int{2121212121}},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(getInvalidIds(test.start, test.end, IsInvalidPart2), test.invalidIds) {
			t.Errorf("GetInvalidIds(%v, %v) is invalid", test.start, test.end)
		}
	}
}

func TestSolveDay2_SolveDay2Part2(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	got := SolveDay2(input, IsInvalidPart2)
	want := 4174379265

	if got != want {
		t.Errorf("No stars today. SolveDay2 Part 2(%v) = %v, want %v", input, got, want)
	}
}

func getInvalidIds(start int, end int, isInvalid func(id int) bool) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= end; i++ {
		if isInvalid(i) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
