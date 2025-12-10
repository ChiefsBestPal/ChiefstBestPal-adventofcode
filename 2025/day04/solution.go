package day04

import (
	"aoc/shared/parser"
	"runtime"
	"sync"
)

type Solution struct{}

type Point struct {
	X, Y int
}

var directions = []Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
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

// findAccessible finds all accessible rolls (simple sequential version)
func findAccessible(grid [][]rune) []Point {
	rows := len(grid)
	cols := len(grid[0])
	var accessible []Point

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '@' {
				if countAdjacentRolls(grid, row, col) < 4 {
					accessible = append(accessible, Point{row, col})
				}
			}
		}
	}

	return accessible
}

// findAccessibleConcurrent finds all accessible rolls in parallel
func findAccessibleConcurrent(grid [][]rune) []Point {
	rows := len(grid)
	cols := len(grid[0])
	numWorkers := runtime.NumCPU()

	// Buffered Channel to collect results asynchronously from workers
	resultCh := make(chan []Point, numWorkers)

	// Divide rows among workers
	rowsPerWorker := (rows + numWorkers - 1) / numWorkers

	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		startRow := w * rowsPerWorker
		endRow := min((w+1)*rowsPerWorker, rows)

		go func(start, end int) {
			defer wg.Done()
			var accessible []Point

			for row := start; row < end; row++ {
				for col := 0; col < cols; col++ {
					if grid[row][col] == '@' {
						if countAdjacentRolls(grid, row, col) < 4 {
							accessible = append(accessible, Point{row, col})
						}
					}
				}
			}

			resultCh <- accessible
		}(startRow, endRow)
	}

	// Close channel when all workers done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect all results
	var allAccessible []Point
	for accessible := range resultCh {
		allAccessible = append(allAccessible, accessible...)
	}

	return allAccessible
}

func (Solution) Part2(input string) any {
	grid := Parse(input)
	totalRemoved := 0

	for {
		// Find all accessible rolls in parallel
		accessible := findAccessible(grid)

		if len(accessible) == 0 {
			break
		}

		// Remove all accessible rolls
		for _, p := range accessible {
			grid[p.X][p.Y] = '.'
		}

		totalRemoved += len(accessible)
	}

	return totalRemoved
}
