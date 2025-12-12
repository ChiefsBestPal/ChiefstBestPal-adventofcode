package day08

import (
	"aoc/shared/parser"
	"math"
	"slices"
)

type Solution struct{}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) DistanceTo(other Point3D) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	dz := float64(p.Z - other.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type Edge struct {
	A, B     int
	Distance float64
}

// UnionFind data structure for tracking connected components
type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already connected
	}

	// Union by size
	if uf.size[rootX] < uf.size[rootY] {
		uf.parent[rootX] = rootY
		uf.size[rootY] += uf.size[rootX]
	} else {
		uf.parent[rootY] = rootX
		uf.size[rootX] += uf.size[rootY]
	}

	return true
}

func (uf *UnionFind) ComponentSizes() []int {
	components := make(map[int]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		components[root] = uf.size[root]
	}

	sizes := make([]int, 0, len(components))
	for _, size := range components {
		sizes = append(sizes, size)
	}
	return sizes
}

func Parse(input string) []Point3D {
	lines := parser.Lines(input)
	points := make([]Point3D, len(lines))

	for i, line := range lines {
		nums := parser.Ints(line)
		points[i] = Point3D{X: nums[0], Y: nums[1], Z: nums[2]}
	}

	return points
}

func (Solution) Part1(input string) any {
	points := Parse(input)
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
	slices.SortFunc(edges, func(a, b Edge) int {
		if a.Distance < b.Distance {
			return -1
		} else if a.Distance > b.Distance {
			return 1
		}
		return 0
	})

	// Connect the first 1000 shortest pairs
	uf := NewUnionFind(n)
	connections := 0

	for _, edge := range edges {
		//fmt.Printf("Considering edge between point %d and %d with distance %.2f\n", points[edge.A], points[edge.B], edge.Distance)
		// if uf.Union(edge.A, edge.B) {
		// 	connections++
		// 	if connections == 1000 {
		// 		break
		// 	}
		// }
		uf.Union(edge.A, edge.B)
		connections++
		if connections == 1000 {
			break
		}
	}
	//fmt.Println("Total connections made:", connections)
	//fmt.Printf("Component sizes: %v\n", uf.size)
	// Get component sizes and find three largest
	sizes := uf.ComponentSizes()
	slices.Sort(sizes)
	slices.Reverse(sizes)
	//fmt.Printf("Component sizes: %v\n", sizes)
	// Multiply three largest
	result := sizes[0]

	if len(sizes) >= 2 && sizes[1] > 0 {
		result *= sizes[1]
	}
	if len(sizes) >= 3 && sizes[2] > 0 {
		result *= sizes[2]
	}

	return result
}

func (Solution) Part2(input string) any {
	return 0 // TODO
}
