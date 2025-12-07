package day04

import (
	"aoc/shared/parser"
)

type Solution struct{}

type Point struct {
	X, Y int
}

// Parse converts input into a 2D grid
func Parse(input string) [][]rune {
	lines := parser.Lines(input)
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

// countAdjacentRolls counts how many '@' symbols are in the 8 adjacent positions
func countAdjacentRolls(grid [][]rune, row, col int) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// 8 directions: N, NE, E, SE, S, SW, W, NW
	directions := []Point{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1},
		{1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}

	for _, dir := range directions {
		newRow := row + dir.X
		newCol := col + dir.Y

		// Check bounds
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

func (Solution) Part1(input string) any {
	grid := Parse(input)
	accessible := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' {
				// A roll is accessible if fewer than 4 adjacent rolls
				adjacentCount := countAdjacentRolls(grid, row, col)
				if adjacentCount < 4 {
					accessible++
				}
			}
		}
	}

	return accessible
}

func (Solution) Part2(input string) any {
	return 0
}
