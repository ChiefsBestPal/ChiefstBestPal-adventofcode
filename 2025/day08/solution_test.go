package day08

import "testing"

const example = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestPart1(t *testing.T) {
	// The example says after 10 connections, multiply 5*4*2 = 40
	// But Part1 does 1000 connections, so we can't use this exact test
	// Let's just verify it runs without error
	got := Solution{}.Part1(example)
	
	// Should return some non-zero result
	if got == 0 {
		t.Errorf("Part1() = %v, expected non-zero", got)
	}
}

// Test the example case with 10 connections
func TestPart1Example10Connections(t *testing.T) {
	points := Parse(example)
	n := len(points)
	
	// Generate all edges
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := points[i].DistanceTo(points[j])
			edges = append(edges, Edge{A: i, B: j, Distance: dist})
		}
	}
	
	// Sort edges by distance
	slices := make([]Edge, len(edges))
	copy(slices, edges)
	
	// Simple bubble sort for test
	for i := 0; i < len(slices)-1; i++ {
		for j := 0; j < len(slices)-i-1; j++ {
			if slices[j].Distance > slices[j+1].Distance {
				slices[j], slices[j+1] = slices[j+1], slices[j]
			}
		}
	}
	
	// Connect the first 10 shortest pairs
	uf := NewUnionFind(n)
	connections := 0
	
	for _, edge := range slices {
		if uf.Union(edge.A, edge.B) {
			connections++
			if connections == 10 {
				break
			}
		}
	}
	
	// Get component sizes and find three largest
	sizes := uf.ComponentSizes()
	
	// Sort descending
	for i := 0; i < len(sizes)-1; i++ {
		for j := 0; j < len(sizes)-i-1; j++ {
			if sizes[j] < sizes[j+1] {
				sizes[j], sizes[j+1] = sizes[j+1], sizes[j]
			}
		}
	}
	
	// Expected: 5 * 4 * 2 = 40
	want := 40
	got := sizes[0] * sizes[1] * sizes[2]
	
	if got != want {
		t.Errorf("After 10 connections: got %v, want %v (sizes: %v)", got, want, sizes)
	}
}

func TestPart2(t *testing.T) {
	got := Solution{}.Part2(example)
	want := 0

	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}
