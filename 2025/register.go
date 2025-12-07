// Package y2025 registers all solutions for Advent of Code 2025.
package y2025

import (
	"aoc/2025/day01"
	"aoc/2025/day02"
	"aoc/2025/day03"
	"aoc/2025/day04"
	"aoc/aoc"
)

func init() {
	aoc.Register(2025, 1, day01.Solution{})
	aoc.Register(2025, 2, day02.Solution{})
	aoc.Register(2025, 3, day03.Solution{})
	aoc.Register(2025, 4, day04.Solution{})

}
