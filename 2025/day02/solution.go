package day02

import (
	"strconv"
	"strings"
)

type Solution struct{}

type Range struct {
	Start int
	End   int
}

func Parse(input string) []Range {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	var ranges []Range
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		bounds := strings.Split(part, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{Start: start, End: end})
	}

	return ranges
}

// isInvalid checks if a product ID is invalid (digits repeated twice)
func isInvalid(id int) bool {
	s := strconv.Itoa(id)

	// Must have even number of digits
	if len(s)%2 != 0 {
		return false
	}

	// First half must equal second half
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}

func (Solution) Part1(input string) any {
	ranges := Parse(input)

	sum := 0
	for _, r := range ranges {
		for id := r.Start; id <= r.End; id++ {
			if isInvalid(id) {
				sum += id
			}
		}
	}

	return sum
}

func (Solution) Part2(input string) any {
	// TODO: implement when revealed
	return 0
}
