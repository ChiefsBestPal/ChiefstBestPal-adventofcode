package day06

import (
	"strconv"
	"strings"
)

type Solution struct{}

type Problem struct {
	Numbers  []int
	Operator rune // '+' or '*'
}

func Parse(input string, extractProblem func([]string, int) *Problem) []Problem {
	//!NB: Don't use parser.Lines() as it trims each line, losing column alignment
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return nil
	}

	// Grid width is zero-padded i.e. it is max line length
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	// Pad all lines
	grid := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < width {
			grid[i] = line + strings.Repeat(" ", width-len(line))
		} else {
			grid[i] = line
		}
	}

	// Extract problems by reading columns
	var problems []Problem
	col := 0

	for col < width {

		if isEmptyColumn(grid, col) {
			col++
			continue
		}

		// Found
		problem := extractProblem(grid, col)
		if problem != nil {
			problems = append(problems, *problem)
		}

		// Skip
		col++
		for col < width && !isEmptyColumn(grid, col) {
			col++
		}
	}

	return problems
}

func isEmptyColumn(grid []string, col int) bool {
	for _, line := range grid {
		if col < len(line) && line[col] != ' ' {
			return false
		}
	}
	return true
}

func (p Problem) Solve() uint64 {
	if len(p.Numbers) == 0 {
		return 0
	}

	result := uint64(p.Numbers[0])

	if p.Operator == '+' {
		for i := 1; i < len(p.Numbers); i++ {
			result += uint64(p.Numbers[i])
		}
	} else if p.Operator == '*' {
		for i := 1; i < len(p.Numbers); i++ {
			result *= uint64(p.Numbers[i])
		}
	}

	return result
}
func extractProblemPart1(grid []string, startCol int) *Problem {
	// Find num of cols (until next empty column or end)
	endCol := startCol
	width := len(grid[0])
	for endCol < width && !isEmptyColumn(grid, endCol) {
		endCol++
	}

	// Extract the column nums + last row operator
	var numbers []int
	var operator rune

	for row := 0; row < len(grid); row++ {
		// Extract text from this row for this problem
		text := strings.TrimSpace(grid[row][startCol:endCol])

		if text == "" {
			continue
		}

		if text == "+" || text == "*" {
			operator = rune(text[0])
		} else {
			// Try to parse as number
			if num, err := strconv.Atoi(text); err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 || operator == 0 {
		return nil
	}

	return &Problem{
		Numbers:  numbers,
		Operator: operator,
	}
}
func (Solution) Part1(input string) any {
	problems := Parse(input, extractProblemPart1)

	var grandTotal uint64
	for _, problem := range problems {
		grandTotal += problem.Solve()
	}

	return grandTotal
}

func extractProblemPart2(grid []string, startCol int) *Problem {
	// Find width of this problem
	endCol := startCol
	width := len(grid[0])
	for endCol < width && !isEmptyColumn(grid, endCol) {
		endCol++
	}

	// Get operator from last row
	var operator rune
	for col := startCol; col < endCol; col++ {
		ch := rune(grid[len(grid)-1][col])
		if ch == '+' || ch == '*' {
			operator = ch
			break
		}
	}

	// Read columns right-to-left, each column forms one number
	// Within each column: top = most significant digit, bottom = least significant
	var numbers []int

	for col := endCol - 1; col >= startCol; col-- {
		// Read this column top-to-bottom to form a number
		numStr := ""
		for row := 0; row < len(grid)-1; row++ { // -1 to skip operator row
			ch := grid[row][col]
			if ch != ' ' {
				numStr += string(ch)
			}
		}

		if numStr != "" {
			if num, err := strconv.Atoi(numStr); err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	if len(numbers) == 0 || operator == 0 {
		return nil
	}

	return &Problem{
		Numbers:  numbers,
		Operator: operator,
	}
}
func (Solution) Part2(input string) any {
	problems := Parse(input, extractProblemPart2)

	var grandTotal uint64
	for _, problem := range problems {
		grandTotal += problem.Solve()
	}

	return grandTotal
}
