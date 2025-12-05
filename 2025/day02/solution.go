package day02

import (
	"strconv"
	"strings"
)

type Solution struct{}

type Range struct {
	Lo, Hi int
}

func Parse(input string) []Range {
	var ranges []Range

	line := strings.TrimSpace(input)
	line = strings.TrimSuffix(line, ",")

	for _, part := range strings.Split(line, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		bounds := strings.Split(part, "-")
		lo, _ := strconv.Atoi(bounds[0])
		hi, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{Lo: lo, Hi: hi})
	}

	return ranges
}

// isDoubled checks if a number is a pattern repeated twice (e.g., 1212, 5555, 123123)
func isDoubled(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	return s[:mid] == s[mid:]
}

// generateDoubled generates all doubled numbers up to maxVal
// by iterating patterns: 1->11, 2->22, ..., 10->1010, 11->1111, etc.
func generateDoubled(maxVal int) []int {
	var result []int

	for pattern := 1; ; pattern++ {
		ps := strconv.Itoa(pattern)
		doubled, _ := strconv.Atoi(ps + ps)

		if doubled > maxVal {
			break
		}
		result = append(result, doubled)
	}

	return result
}

func (Solution) Part1(input string) any {
	ranges := Parse(input)

	// Find max value across all ranges
	maxVal := 0
	for _, r := range ranges {
		if r.Hi > maxVal {
			maxVal = r.Hi
		}
	}

	// Generate all possible doubled numbers up to max
	doubled := generateDoubled(maxVal)

	// Sum those that fall in any range
	sum := 0
	for _, d := range doubled {
		for _, r := range ranges {
			if d >= r.Lo && d <= r.Hi {
				sum += d
				break // Don't double-count if in multiple ranges
			}
		}
	}

	return sum
}

func (Solution) Part2(input string) any {
	// TODO: implement when revealed
	return 0
}
