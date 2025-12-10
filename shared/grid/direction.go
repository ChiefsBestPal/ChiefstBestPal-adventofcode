package grid

// Direction represents a cardinal direction
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// All4Directions returns all 4 cardinal directions
var All4Directions = []Direction{North, East, South, West}

// Vector returns the unit vector for this direction
func (d Direction) Vector() Point {
	switch d {
	case North:
		return Point{0, -1}
	case East:
		return Point{1, 0}
	case South:
		return Point{0, 1}
	case West:
		return Point{-1, 0}
	default:
		return Point{0, 0}
	}
}

// TurnLeft returns the direction 90 degrees counter-clockwise
func (d Direction) TurnLeft() Direction {
	return Direction((int(d) + 3) % 4)
}

// TurnRight returns the direction 90 degrees clockwise
func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

// Reverse returns the opposite direction
func (d Direction) Reverse() Direction {
	return Direction((int(d) + 2) % 4)
}

// String returns the name of the direction
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

// FromRune converts a rune to a Direction (^>v< or NESW)
func FromRune(r rune) (Direction, bool) {
	switch r {
	case '^', 'N', 'U':
		return North, true
	case '>', 'E', 'R':
		return East, true
	case 'v', 'V', 'S', 'D':
		return South, true
	case '<', 'W', 'L':
		return West, true
	default:
		return North, false
	}
}
