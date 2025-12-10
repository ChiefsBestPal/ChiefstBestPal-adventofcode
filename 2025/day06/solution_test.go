package day06

import "testing"

const example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 4277556

	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Solution{}.Part2(example)
	want := "TODO"

	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
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
