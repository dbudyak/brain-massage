package main

import (
	"reflect"
	"testing"
)

func TestSolveDay2_IsInvlaid(t *testing.T) {
	valid := 12312

	if IsInvalid(valid) {
		t.Errorf("%v is invalid", valid)
	}

	valid = 1

	if IsInvalid(valid) {
		t.Errorf("%v is invalid", valid)
	}

	invalid := 1212
	if !IsInvalid(invalid) {
		t.Errorf("%v is valid", valid)
	}
}

func TestSolveDay2_GetInvalidIds(t *testing.T) {
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
		if !reflect.DeepEqual(GetInvalidIds(test.start, test.end), test.invalidIds) {
			t.Errorf("GetInvalidIds(%v, %v) is invalid", test.start, test.end)
		}
	}
}

func TestSolveDay2_SolveDay2Part1(t *testing.T) {
	input := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	got := SolveDay2Part1(input)
	want := 1227775554

	if got != want {
		t.Errorf("No stars today. SolveDay2Part1(%v) = %v, want %v", input, got, want)
	}
}
