package day02

import "testing"

const example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

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

func TestIsDoubled(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{11, true},
		{22, true},
		{55, true},
		{99, true},
		{1010, true},
		{6464, true},
		{123123, true},
		{222222, true},
		{1188511885, true},
		{12, false},
		{101, false},   // odd length
		{1234, false},  // 12 != 34
		{12345, false}, // odd length
	}

	for _, tt := range tests {
		got := isDoubled(tt.n)
		if got != tt.want {
			t.Errorf("isDoubled(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestGenerateDoubled(t *testing.T) {
	doubled := generateDoubled(100)
	// Should be: 11, 22, 33, 44, 55, 66, 77, 88, 99
	want := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	if len(doubled) != len(want) {
		t.Errorf("generateDoubled(100) got %d items, want %d", len(doubled), len(want))
		return
	}

	for i, v := range want {
		if doubled[i] != v {
			t.Errorf("generateDoubled(100)[%d] = %d, want %d", i, doubled[i], v)
		}
	}
}

func TestParse(t *testing.T) {
	ranges := Parse("11-22,95-115")

	if len(ranges) != 2 {
		t.Fatalf("Parse() got %d ranges, want 2", len(ranges))
	}

	if ranges[0].Lo != 11 || ranges[0].Hi != 22 {
		t.Errorf("ranges[0] = %v, want {11, 22}", ranges[0])
	}

	if ranges[1].Lo != 95 || ranges[1].Hi != 115 {
		t.Errorf("ranges[1] = %v, want {95, 115}", ranges[1])
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
