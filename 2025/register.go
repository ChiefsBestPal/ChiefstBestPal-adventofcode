// Package y2025 registers all solutions for Advent of Code 2025.
package y2025

import (
	"aoc/2025/day01"
	"aoc/cmd/aoc"
)

func init() {
	main.Register(2025, 1, day01.Solution{})
	// Add more days as you solve them:
	// main.Register(2025, 2, day02.Solution{})
}
