package day02

import (
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

type Range struct {
	Lo, Hi int
}

// Current optimization: Cache with hashed values and keys as binary search 'lst[i].Hi >= inserted.Lo'
// TODO: Improve it using interval or segment trees to optimize BST

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

	// Sort ranges by Lo for binary search optimization
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Lo < ranges[j].Lo
	})

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

	// Generate all possible doubled numbers up to max (already sorted)
	doubled := generateDoubled(maxVal)

	sum := 0
	rangeIdx := 0 // Two-pointer approach: both doubled and ranges are sorted

	// For each doubled number, check ranges using two-pointer technique
	// This avoids repeated binary searches and achieves O(n + m) instead of O(n*log(m))
	for _, d := range doubled {
		// Skip ranges that are completely below d (optimization: pointer never goes back)
		for rangeIdx < len(ranges) && ranges[rangeIdx].Hi < d {
			rangeIdx++
		}

		// Check ranges that could contain d
		for i := rangeIdx; i < len(ranges); i++ {
			r := ranges[i]
			// If d is before this range, it won't be in any later ranges
			if d < r.Lo {
				break
			}
			// Check if d is in this range
			if d <= r.Hi {
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
