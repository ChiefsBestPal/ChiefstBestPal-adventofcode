package day01

import (
	"aoc/shared/math"
	"aoc/shared/parser"
	"strconv"
)

type Solution struct{}

type Rotation struct {
	Dir  int // -1 = Left, +1 = Right
	Dist int
}

func Parse(input string) []Rotation {
	var rotations []Rotation

	for _, line := range parser.Lines(input) {

		dir := 1
		if line[0] == 'L' {
			dir = -1
		}

		dist, _ := strconv.Atoi(line[1:])
		rotations = append(rotations, Rotation{Dir: dir, Dist: dist})
	}

	return rotations
}

func (Solution) Part1(input string) any {
	rotations := Parse(input)

	pos := 50
	count := 0

	for _, r := range rotations {
		pos = math.Mod(pos+r.Dir*r.Dist, 100)
		if pos == 0 {
			count++
		}
	}

	return count
}

func (Solution) Part2(input string) any {
	rotations := Parse(input)

	pos := 50
	count := 0
	var remainder int
	for _, r := range rotations {
		remainder = r.Dist - (r.Dist/100)*100
		if pos != 0 && ((r.Dir == -1 && pos <= remainder) || (r.Dir == 1 && (100-pos) <= remainder)) {
			count++
		}

		count += r.Dist / 100 // Full turns

		pos = math.Mod(pos+r.Dir*r.Dist, 100)
	}

	return count
}
