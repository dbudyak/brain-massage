package main

import "testing"

const input9 = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestDay9Part1(t *testing.T) {
	got := SolveDay9Part1(input9)
	want := 50
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
