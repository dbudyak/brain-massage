package main

import (
	"testing"
)

func TestSolveDay1(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	want := 3
	got := SolveDay1(input)

	if got != want {
		t.Errorf("SolveDay1() = %d, want %d", got, want)
	}
}

func TestSolveDay1Part2(t *testing.T) {
	input := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	want := 6
	got := SolveDay1Part2(input)

	if got != want {
		t.Errorf("SolveDay1Part2() = %d, want %d", got, want)
	}
}

func TestCountZeroCrossings(t *testing.T) {
	tests := []struct {
		name      string
		position  int
		distance  int
		direction int
		want      int
	}{
		{"L68 from 50", 50, 68, left, 1},
		{"L30 from 82", 82, 30, left, 0},
		{"R48 from 52", 52, 48, right, 1},
		{"L5 from 0", 0, 5, left, 0},
		{"R60 from 95", 95, 60, right, 1},
		{"L55 from 55", 55, 55, left, 1},
		{"L1 from 0", 0, 1, left, 0},
		{"L99 from 99", 99, 99, left, 1},
		{"R14 from 0", 0, 14, right, 0},
		{"L82 from 14", 14, 82, left, 1},
		{"R1000 from 50", 50, 1000, right, 10},
		{"L250 from 30", 30, 250, left, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountZeroCrossings(tt.position, tt.distance, tt.direction)
			if got != tt.want {
				t.Errorf("CountZeroCrossings(%d, %d, %d) = %d, want %d",
					tt.position, tt.distance, tt.direction, got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name      string
		position  int
		distance  int
		direction int
		want      int
	}{
		{"R8 from 11", 11, 8, right, 19},
		{"L19 from 19", 19, 19, left, 0},
		{"L10 from 5", 5, 10, left, 95},
		{"R5 from 95", 95, 5, right, 0},
		{"L68 from 50", 50, 68, left, 82},
		{"wrap right from 99", 99, 1, right, 0},
		{"wrap left from 0", 0, 1, left, 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rotate(tt.position, tt.distance, tt.direction)
			if got != tt.want {
				t.Errorf("Rotate(%d, %d, %d) = %d, want %d",
					tt.position, tt.distance, tt.direction, got, tt.want)
			}
		})
	}
}
