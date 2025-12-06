package main

import "testing"

func TestSolveDay6Part1(t *testing.T) {
	input := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "

	got := SolveDay6Part1(input)
	want := 4277556

	if got != want {
		t.Errorf("SolveDay6Part1() = %d, want %d", got, want)
	}
}

/*func TestSolveDay6Part2(t *testing.T) {
	input := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "

	got := SolveDay6Part2(input)
	want := int64(3263827)

	if got != want {
		t.Errorf("SolveDay6Part2() = %d, want %d", got, want)
	}
}*/

/*func TestSpecialMath(t *testing.T) {
	got := DoSpecialMath([]string{"64", "23", "314"}, "+")
	if got != 1058 {
		t.Errorf("DoSpecialMath() = %d, want 1058", got)
	}
}
*/
