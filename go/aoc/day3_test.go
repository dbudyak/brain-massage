package main

import "testing"

func TestSolveDay3Part1(t *testing.T) {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	got := SolveDay3Part1(input)
	want := 357

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSolveDay3Part1_GetMaximumJoltage(t *testing.T) {
	tests := []struct {
		battery    string
		maxJoltage int
	}{
		{battery: "987654321111111", maxJoltage: 98},
		{battery: "811111111111119", maxJoltage: 89},
		{battery: "234234234234278", maxJoltage: 78},
		{battery: "818181911112111", maxJoltage: 92},
	}

	for _, test := range tests {
		got := GetMaximumJoltage(test.battery)
		want := test.maxJoltage
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestSolveDay3Part2(t *testing.T) {
	input := "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	got := SolveDay3Part2(input)
	want := 3121910778619

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSolveDay3Part2_GetMaximumJoltage(t *testing.T) {
	tests := []struct {
		battery    string
		maxJoltage int
	}{
		{battery: "987654321111111", maxJoltage: 987654321111},
		{battery: "811111111111119", maxJoltage: 811111111119},
		{battery: "234234234234278", maxJoltage: 434234234278},
		{battery: "818181911112111", maxJoltage: 888911112111},
	}

	for _, test := range tests {
		got := GetMaximumJoltagePart2(test.battery, 12)
		want := test.maxJoltage
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
