package grid

import (
	"aoc/shared/parser"
)

type Grid[T any] struct {
	Data   [][]T
	Width  int
	Height int
}

func NewGrid[T any](width, height int) *Grid[T] {
	data := make([][]T, height)
	for i := range data {
		data[i] = make([]T, width)
	}
	return &Grid[T]{
		Data:   data,
		Width:  width,
		Height: height,
	}
}

// NewGridFromLines creates a grid from lines of text, applying a converter function
func NewGridFromLines[T any](input string, convert func(rune) T) *Grid[T] {
	lines := parser.Lines(input)
	if len(lines) == 0 {
		return &Grid[T]{}
	}

	height := len(lines)
	width := len(lines[0])
	g := NewGrid[T](width, height)

	for y, line := range lines {
		for x, ch := range line {
			g.Data[y][x] = convert(ch)
		}
	}

	return g
}

// func NewRuneGrid(input string) *Grid[rune] {
// 	return NewGridFromLines(input, func(r rune) rune { return r })
// }

//	func NewByteGrid(input string) *Grid[byte] {
//		return NewGridFromLines(input, func(r rune) byte { return byte(r) })
//	}
//
// // String converts a rune grid to a string representation
//
//	func (g *Grid[rune]) String() string {
//		var sb strings.Builder
//		for y := 0; y < g.Height; y++ {
//			for x := 0; x < g.Width; x++ {
//				sb.WriteRune(g.Data[y][x])
//			}
//			if y < g.Height-1 {
//				sb.WriteRune('\n')
//			}
//		}
//		return sb.String()
//	}
//
// InBounds checks if a point is within the grid boundaries
func (g *Grid[T]) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.Width && p.Y >= 0 && p.Y < g.Height
}

func (g *Grid[T]) Get(p Point) T {
	if !g.InBounds(p) {
		var zero T
		return zero
	}
	return g.Data[p.Y][p.X]
}

func (g *Grid[T]) Set(p Point, val T) {
	if g.InBounds(p) {
		g.Data[p.Y][p.X] = val
	}
}

// Find returns the first (any) point where the predicate is true
func (g *Grid[T]) Find(predicate func(T) bool) (Point, bool) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if predicate(g.Data[y][x]) {
				return Point{x, y}, true
			}
		}
	}
	return Point{}, false
}

// FindAll returns all points where the predicate is true
func (g *Grid[T]) FindAll(predicate func(T) bool) []Point {
	var points []Point
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if predicate(g.Data[y][x]) {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

// Count counts cells matching the predicate
func (g *Grid[T]) Count(predicate func(T) bool) int {
	count := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if predicate(g.Data[y][x]) {
				count++
			}
		}
	}
	return count
}

// ForEach applies a function to each cell
func (g *Grid[T]) ForEach(fn func(p Point, val T)) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			fn(Point{x, y}, g.Data[y][x])
		}
	}
}

// Clone creates a deep copy of the grid
func (g *Grid[T]) Clone() *Grid[T] {
	clone := NewGrid[T](g.Width, g.Height)
	for y := 0; y < g.Height; y++ {
		copy(clone.Data[y], g.Data[y])
	}
	return clone
}

// Neighbors4 returns valid 4-directional neighbors of a point
func (g *Grid[T]) Neighbors4(p Point) []Point {
	var neighbors []Point
	for _, n := range p.Neighbors4() {
		if g.InBounds(n) {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// Neighbors8 returns valid 8-directional neighbors of a point
func (g *Grid[T]) Neighbors8(p Point) []Point {
	var neighbors []Point
	for _, n := range p.Neighbors8() {
		if g.InBounds(n) {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}
