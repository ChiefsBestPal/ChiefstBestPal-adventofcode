package grid

import "math"

// Point represents a 2D coordinate in a grid, cartesian plane or matrix
type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p Point) Mul(scalar int) Point {
	return Point{X: p.X * scalar, Y: p.Y * scalar}
}

func (p Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func (p Point) ChebyshevDistance(other Point) int {
	return max(abs(p.X-other.X), abs(p.Y-other.Y))
}

func (p Point) EuclideanDistance(other Point) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) Rotate90() Point {
	return Point{X: p.Y, Y: -p.X}
}

func (p Point) Rotate180() Point {
	return Point{X: -p.X, Y: -p.Y}
}

func (p Point) Rotate270() Point {
	return Point{X: -p.Y, Y: p.X}
}

func (p Point) Neighbors4() []Point {
	return []Point{
		{p.X, p.Y - 1}, // North
		{p.X + 1, p.Y}, // East
		{p.X, p.Y + 1}, // South
		{p.X - 1, p.Y}, // West
	}
}

func (p Point) Neighbors8() []Point {
	return []Point{
		{p.X, p.Y - 1},     // N
		{p.X + 1, p.Y - 1}, // NE
		{p.X + 1, p.Y},     // E
		{p.X + 1, p.Y + 1}, // SE
		{p.X, p.Y + 1},     // S
		{p.X - 1, p.Y + 1}, // SW
		{p.X - 1, p.Y},     // W
		{p.X - 1, p.Y - 1}, // NW
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
