package day05

import (
	interval "aoc/shared/intervals"
	"aoc/shared/parser"
	"strconv"
	"strings"
)

type Solution struct{}

type Range = interval.Range

type Database struct {
	FreshRanges   []Range
	IngredientIDs []int
}

func Parse(input string) Database {
	sections := strings.Split(input, "\n\n")

	var db Database

	// Parse fresh ranges
	for _, line := range parser.Lines(sections[0]) {
		parts := strings.Split(line, "-")
		lo, _ := strconv.Atoi(parts[0])
		hi, _ := strconv.Atoi(parts[1])
		db.FreshRanges = append(db.FreshRanges, Range{Lo: lo, Hi: hi})
	}

	// Parse ingredient IDs
	for _, line := range parser.Lines(sections[1]) {
		id, _ := strconv.Atoi(line)
		db.IngredientIDs = append(db.IngredientIDs, id)
	}

	return db
}

func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.Lo && id <= r.Hi {
			return true
		}
	}
	return false
}

func (Solution) Part1(input string) any {
	db := Parse(input)

	freshCount := 0
	for _, id := range db.IngredientIDs {
		if isFresh(id, db.FreshRanges) {
			freshCount++
		}
	}

	return freshCount
}

func (Solution) Part2(input string) any {
	db := Parse(input)
	count := 0

	for _, r := range interval.MergeUnsorted(db.FreshRanges) {
		count += r.Size()
	}
	return count
}
