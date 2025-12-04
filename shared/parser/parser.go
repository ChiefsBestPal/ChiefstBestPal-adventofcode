package parser

import (
	"strconv"
	"strings"
)

// Lines splits input into non-empty trimmed lines
func Lines(input string) []string {
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// Sections splits input by blank lines (common in AOC)
func Sections(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}

// Ints extracts all integers from a string (handles negative numbers)
func Ints(s string) []int {
	var nums []int
	var current strings.Builder
	inNumber := false
	isNegative := false

	for i, ch := range s {
		if ch == '-' && (i == 0 || !inNumber) && i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
			isNegative = true
			inNumber = true
		} else if ch >= '0' && ch <= '9' {
			if !inNumber {
				inNumber = true
			}
			current.WriteRune(ch)
		} else if inNumber {
			numStr := current.String()
			if n, err := strconv.Atoi(numStr); err == nil {
				if isNegative {
					n = -n
				}
				nums = append(nums, n)
			}
			current.Reset()
			inNumber = false
			isNegative = false
		}
	}

	// Handle last number
	if inNumber {
		numStr := current.String()
		if n, err := strconv.Atoi(numStr); err == nil {
			if isNegative {
				n = -n
			}
			nums = append(nums, n)
		}
	}

	return nums
}

// Grid parses input into a 2D character grid
func Grid(input string) [][]rune {
	lines := Lines(input)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}
