// Package y2025 registers all solutions for Advent of Code 2025.
package y2025

import (
	"aoc/2025/day01"
	"aoc/aoc"
)

func init() {
	aoc.Register(2025, 1, day01.Solution{})
	// Add more days as you solve them:
	// aoc.Register(2025, 2, day02.Solution{})
}
