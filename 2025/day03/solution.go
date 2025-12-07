package day03

import (
	"aoc/shared/parser"
	"math"
)

type Solution struct{}

// maxJoltage finds the maximum joltage for a single bank
// by picking N batteries (maintaining their order) that form the largest number
func maxJoltage(bank string, N int) int {
	if len(bank) < N {
		return 0
	}

	jolt := make([]int, N)
	jolt[0] = int(bank[0] - '0')

	for i := 1; i < len(bank); i++ {
		dig := int(bank[i] - '0')
		remainingLen := len(bank) - i - 1
		for j := 0; j < N; j++ {
			if jolt[j] < dig && remainingLen >= (N-j-1) {
				jolt[j] = dig
				// jolt = jolt[:j+1] // Logical length reduced to position of digit just updated
				// printSlice(jolt)
				clear(jolt[j+1 : N]) // turn all other underlying digits after current back to zero

				break
			}
		}
	}
	result := 0
	for pos, dig := range jolt {
		result += int(math.Pow10(N-pos-1)) * dig
	}
	return result
}

func (Solution) Part1(input string) any {
	lines := parser.Lines(input)

	total := 0
	for _, line := range lines {
		total += maxJoltage(line, 2)
	}

	return total
}

func (Solution) Part2(input string) any {
	lines := parser.Lines(input)
	total := 0
	for _, line := range lines {
		total += maxJoltage(line, 12)
	}

	return total
}
