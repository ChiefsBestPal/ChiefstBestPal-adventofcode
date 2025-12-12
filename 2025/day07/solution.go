package day07

import (
	"aoc/shared/grid"
)

type Solution struct{}

type Beam struct {
	Row, Col int
	Dir      grid.Direction // Direction beam is traveling
}

func Parse(input string) *grid.Grid[rune] {
	return grid.NewGridFromLines(input, func(r rune) rune { return r })
}

func (Solution) Part1(input string) any {
	g := Parse(input)

	// Find starting position 'S'
	start, found := g.Find(func(r rune) bool { return r == 'S' })
	if !found {
		return 0
	}
	
	// Track beams and splits
	splitCount := 0
	visited := make(map[Beam]bool) // Prevent infinite loops
	
	// BFS queue of beams (Point is X=col, Y=row in grid coordinates)
	queue := []Beam{{Row: start.Y, Col: start.X, Dir: grid.South}}

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		// Check if already visited this beam state
		if visited[beam] {
			continue
		}
		visited[beam] = true

		// Move beam in its direction
		next := beam
		switch beam.Dir {
		case grid.North:
			next.Row--
		case grid.South:
			next.Row++
		case grid.East:
			next.Col++
		case grid.West:
			next.Col--
		}

		// Check if beam exits manifold
		if !g.InBounds(grid.Point{X: next.Col, Y: next.Row}) {
			continue
		}

		cell := g.Get(grid.Point{X: next.Col, Y: next.Row})
		
		// If empty space, continue beam
		if cell == '.' || cell == 'S' {
			queue = append(queue, next)
			continue
		}
		
		// If splitter, create two new beams
		if cell == '^' {
			splitCount++
			// Create left and right beams from the splitter
			leftBeam := Beam{Row: next.Row, Col: next.Col, Dir: grid.West}
			rightBeam := Beam{Row: next.Row, Col: next.Col, Dir: grid.East}
			queue = append(queue, leftBeam, rightBeam)
		}
	}
	
	return splitCount
}

func (Solution) Part2(input string) any {
	return 0 // TODO
}
