package graph

type UnionFind[T comparable] interface {
	Find(x T) T
	Union(a, b T)
	Same(a, b T) bool
	Size(x T) int
}

type UFRank[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	sz     map[T]int // size can still be tracked
}

func NewUFRank[T comparable]() *UFRank[T] {
	return &UFRank[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		sz:     make(map[T]int),
	}
}
func (uf *UFRank[T]) init(x T) {
	if _, ok := uf.parent[x]; !ok {
		uf.parent[x] = x
		uf.rank[x] = 0
		uf.sz[x] = 1
	}
}

func (uf *UFRank[T]) Find(x T) T {
	uf.init(x)
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UFRank[T]) Union(a, b T) {
	ra := uf.Find(a)
	rb := uf.Find(b)
	if ra == rb {
		return
	}

	// union by rank
	switch {
	case uf.rank[ra] < uf.rank[rb]:
		ra, rb = rb, ra
	case uf.rank[ra] == uf.rank[rb]:
		uf.rank[ra]++
	}
	// // union by size
	// if uf.sz[ra] < uf.sz[rb] {
	// 	ra, rb = rb, ra
	// }
	uf.parent[rb] = ra
	uf.sz[ra] += uf.sz[rb]
}

func (uf *UFRank[T]) Same(a, b T) bool {
	return uf.Find(a) == uf.Find(b)
}

func (uf *UFRank[T]) Size(x T) int {
	return uf.sz[uf.Find(x)]
}
