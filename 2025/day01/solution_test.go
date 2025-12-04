package day01

import (
	"aoc/shared/math"
	"testing"
)

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 3

	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Solution{}.Part2(example)
	want := 0

	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		a, m, want int
	}{
		{5, 100, 5},
		{-5, 100, 95},
		{105, 100, 5},
		{0, 100, 0},
	}

	for _, tt := range tests {
		got := math.Mod(tt.a, tt.m)
		if got != tt.want {
			t.Errorf("mod(%d, %d) = %d, want %d", tt.a, tt.m, got, tt.want)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution{}.Part1(example)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution{}.Part2(example)
	}
}
