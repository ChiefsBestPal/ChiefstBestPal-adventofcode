package day02

import (
	interval "aoc/shared/intervals"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}

type AccumulateNode struct {
	Num, Sum int
}

type Range = interval.Range

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
	interval.Sort(ranges)

	return ranges
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

func buildRepeated(pattern, reps int) int {
	s := strconv.Itoa(pattern)
	result := strings.Repeat(s, reps)
	n, _ := strconv.Atoi(result)
	return n
}
func generateRepeated(maxVal int) (result []AccumulateNode) {
	seen := make(map[int]bool)
	result = append(result, AccumulateNode{Num: 0, Sum: 0})
	maxDigits := int(math.Floor(math.Log10(float64(maxVal)))) + 1

	// For each pattern length
	for patternLen := 1; patternLen <= maxDigits; patternLen++ {
		start := int(math.Pow10(patternLen - 1))
		end := start * 10

		// For each number of repetitions 'pattern' can be repeated
		for reps := 2; (patternLen * reps) <= maxDigits; reps++ {

			// Check if smallest possible number exceeds maxVal
			if repeated := buildRepeated(start, reps); repeated > maxVal {
				break
			} else if !seen[repeated] {
				result = append(result, AccumulateNode{Num: repeated})
				seen[repeated] = true
			}

			// Generate all other patterns of this length and reps
			for pattern := start + 1; pattern < end; pattern++ {
				repeated := buildRepeated(pattern, reps)
				if repeated > maxVal {
					break
				} else if !seen[repeated] {
					result = append(result, AccumulateNode{Num: repeated})
					seen[repeated] = true
				}
			}
		}
	}

	// Sort by Num in ascending order then compute accumulative Sum of nums...
	slices.SortFunc(result, func(a, b AccumulateNode) int {
		return a.Num - b.Num
	})

	accSum := 0
	for i := range result {
		accSum += result[i].Num
		result[i].Sum = accSum
	}

	return result
}

func (Solution) Part2(input string) any {
	ranges := Parse(input)
	ranges = interval.Merge(ranges) // Merge overlapping ranges to avoid double-counting

	maxVal := 0
	for _, r := range ranges {
		if r.Hi > maxVal {
			maxVal = r.Hi
		}
	}

	repeated_acc := generateRepeated(maxVal)
	cmpFunc := func(a AccumulateNode, targetNum int) int {
		switch {
		case a.Num < targetNum:
			return -1
		case a.Num > targetNum:
			return 1
		default:
			return 0
		}
	}
	sum := 0
	for _, r := range ranges {
		// Find first element >= r.Lo
		startIx, _ := slices.BinarySearchFunc(repeated_acc, r.Lo, cmpFunc)
		// startIx is correct: exact match or insertion point (first >= r.Lo)

		// Find last element <= r.Hi
		endIx, endFound := slices.BinarySearchFunc(repeated_acc, r.Hi, cmpFunc)
		if !endFound {
			// endIx is insertion point, decrement to get last element < r.Hi
			endIx--
		}
		// If found, endIx is exact match (which we want to include)

		// Calculate range sum: Sum[0..endIx] - Sum[0..startIx-1]
		if startIx <= endIx && startIx < len(repeated_acc) && endIx >= 0 {
			sumBefore := 0
			if startIx > 0 {
				sumBefore = repeated_acc[startIx-1].Sum
			}
			sum += repeated_acc[endIx].Sum - sumBefore
		}

	}

	return sum
}
