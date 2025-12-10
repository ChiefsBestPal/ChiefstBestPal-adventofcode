package grid

import (
	"testing"
)

func TestPoint(t *testing.T) {
	p1 := Point{3, 4}
	p2 := Point{1, 2}

	if got := p1.Add(p2); got != (Point{4, 6}) {
		t.Errorf("Add: got %v, want {4, 6}", got)
	}

	if got := p1.ManhattanDistance(p2); got != 4 {
		t.Errorf("ManhattanDistance: got %d, want 4", got)
	}

	neighbors := p1.Neighbors4()
	if len(neighbors) != 4 {
		t.Errorf("Neighbors4: got %d neighbors, want 4", len(neighbors))
	}
}

func TestDirection(t *testing.T) {
	d := North

	if got := d.TurnRight(); got != East {
		t.Errorf("TurnRight: got %v, want East", got)
	}

	if got := d.TurnLeft(); got != West {
		t.Errorf("TurnLeft: got %v, want West", got)
	}

	if got := d.Reverse(); got != South {
		t.Errorf("Reverse: got %v, want South", got)
	}

	vec := d.Vector()
	if vec != (Point{0, -1}) {
		t.Errorf("Vector: got %v, want {0, -1}", vec)
	}
}
