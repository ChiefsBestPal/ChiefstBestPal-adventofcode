package day02

import (
	"testing"
)

const example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,
1698522-1698528,446443-446449,38593856-38593862,565653-565659,
824824821-824824827,2121212118-2121212124`

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 1227775554

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

func TestIsInvalid(t *testing.T) {
	tests := []struct {
		id   int
		want bool
	}{
		{11, true},     // 1 repeated twice
		{22, true},     // 2 repeated twice
		{55, true},     // 5 repeated twice
		{99, true},     // 9 repeated twice
		{6464, true},   // 64 repeated twice
		{123123, true}, // 123 repeated twice
		{1010, true},   // 10 repeated twice
		{1188511885, true}, // 11885 repeated twice
		{222222, true}, // 222 repeated twice
		{446446, true}, // 446 repeated twice
		{38593859, true}, // 3859 repeated twice
		{12, false},    // different digits
		{101, false},   // odd number of digits
		{1234, false},  // first half != second half
		{1698522, false}, // not a repeated pattern
	}

	for _, tt := range tests {
		got := isInvalid(tt.id)
		if got != tt.want {
			t.Errorf("isInvalid(%d) = %v, want %v", tt.id, got, tt.want)
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
