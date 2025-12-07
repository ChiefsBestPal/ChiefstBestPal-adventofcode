package day03

import (
	"aoc/shared/parser"
)

type Solution struct{}

// maxJoltage finds the maximum joltage for a single bank
// by picking two batteries (maintaining their order) that form the largest number
func maxJoltage(bank string) int {
	if len(bank) < 2 {
		return 0
	}

	maxVal := 0

	// Try all pairs of positions (i, j) where i < j
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			// Form the two-digit number from batteries at positions i and j
			val := int(bank[i]-'0')*10 + int(bank[j]-'0')
			if val > maxVal {
				maxVal = val
			}
		}
	}

	return maxVal
}

func (Solution) Part1(input string) any {
	lines := parser.Lines(input)

	total := 0
	for _, line := range lines {
		total += maxJoltage(line)
	}

	return total
}

func (Solution) Part2(input string) any {
	// TODO: implement when revealed
	return 0
}
