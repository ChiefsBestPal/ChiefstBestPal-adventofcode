package day03

import "testing"

const example = `987654321111111
811111111111119
234234234234278
818181911112111`

const arbitrary_N_batteries = 5

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 357

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

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		bank string
		want int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
		{"12345", 45}, // positions 2(value=3) and 4(value=5) not optimal; best is pos 3(4) and 4(5) = 45
		{"99", 99},
		{"11", 11},
		{"9876543210", 98},
		{"54321", 54}, // positions 0(5) and 1(4) = 54
	}

	for _, tt := range tests {
		got := maxJoltage(tt.bank, arbitrary_N_batteries)
		if got != tt.want {
			t.Errorf("maxJoltage(%q) = %v, want %v", tt.bank, got, tt.want)
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
