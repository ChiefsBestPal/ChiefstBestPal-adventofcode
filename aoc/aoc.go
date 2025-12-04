// Package aoc provides the solution registry for Advent of Code
package aoc

// Solution interface that each day implements
type Solution interface {
	Part1(input string) any
	Part2(input string) any
}

// Solutions maps year -> day -> solution
var Solutions = make(map[int]map[int]Solution)

// Register adds a solution to the registry
func Register(year, day int, sol Solution) {
	if Solutions[year] == nil {
		Solutions[year] = make(map[int]Solution)
	}
	Solutions[year][day] = sol
}
