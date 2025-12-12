package day07

import "aoc/shared/grid"

type Solution struct{}

func Parse(input string) *grid.Grid[rune] {
	return grid.NewGridFromLines(input, func(r rune) rune { return r })
}

func (Solution) Part1(input string) any {
	g := Parse(input)
	start, found := g.Find(func(r rune) bool { return r == 'S' })
	if !found {
		return 0
	}

	splitCount := 0
	visited := make(map[grid.Point]bool)
	queue := []grid.Point{{X: start.X, Y: start.Y + 1}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if visited[p] || !g.InBounds(p) {
			continue
		}
		visited[p] = true

		cell := g.Get(p)
		if cell == '.' || cell == 'S' {
			queue = append(queue, grid.Point{X: p.X, Y: p.Y + 1})
		} else if cell == '^' {
			splitCount++
			queue = append(queue, grid.Point{X: p.X - 1, Y: p.Y + 1})
			queue = append(queue, grid.Point{X: p.X + 1, Y: p.Y + 1})
		}
	}

	return splitCount
}

func (Solution) Part2(input string) any {
	return 0
}
