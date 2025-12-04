package main

import "testing"

func TestSolveDay4Part1(t *testing.T) {
	input := `
		..@@.@@@@.
		@@@.@.@.@@
		@@@@@.@.@@
		@.@@@@..@.
		@@.@@@@.@@
		.@@@@@@@.@
		.@.@.@.@@@
		@.@@@.@@@@
		.@@@@@@@@.
		@.@.@@@.@.
		`

	got := SolveDay4Part1(input)
	want := 13

	if got != want {
		t.Errorf("SolveDay4Part1() = %d, want %d", got, want)
	}

}

func TestSolveDay4Part2(t *testing.T) {
	input := `
		..@@.@@@@.
		@@@.@.@.@@
		@@@@@.@.@@
		@.@@@@..@.
		@@.@@@@.@@
		.@@@@@@@.@
		.@.@.@.@@@
		@.@@@.@@@@
		.@@@@@@@@.
		@.@.@@@.@.
		`

	got := SolveDay4Part2(input)
	want := 43

	if got != want {
		t.Errorf("SolveDay4Part1() = %d, want %d", got, want)
	}

}
