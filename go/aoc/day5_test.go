package main

import "testing"

func TestSolveDay5Part1(t *testing.T) {
	input := "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"

	got := SolveDay5Part1(input)
	want := 3

	if got != want {
		t.Errorf("SolveDay5Part1() = %d, want %d", got, want)
	}

}

func TestSolveDay5Part2(t *testing.T) {
	input := "3-5\n10-14\n16-20\n12-18\n\n1\n5\n8\n11\n17\n32"

	got := SolveDay5Part2(input)
	want := 14

	if got != want {
		t.Errorf("SolveDay5Part1() = %d, want %d", got, want)
	}

}
