package grid

import "container/heap"

// BFS performs breadth-first search from start to goal
// Returns the path and distance, or nil and -1 if no path exists
func BFS[T any](g *Grid[T], start, goal Point, walkable func(T) bool) ([]Point, int) {
	if !g.InBounds(start) || !g.InBounds(goal) {
		return nil, -1
	}

	queue := []Point{start}
	visited := make(map[Point]bool)
	parent := make(map[Point]Point)
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == goal {
			// Reconstruct path
			path := []Point{}
			for p := goal; p != start; p = parent[p] {
				path = append([]Point{p}, path...)
			}
			path = append([]Point{start}, path...)
			return path, len(path) - 1
		}

		for _, next := range g.Neighbors4(current) {
			if !visited[next] && walkable(g.Get(next)) {
				visited[next] = true
				parent[next] = current
				queue = append(queue, next)
			}
		}
	}

	return nil, -1
}

// FloodFill finds all reachable points from start
func FloodFill[T any](g *Grid[T], start Point, walkable func(T) bool) map[Point]int {
	if !g.InBounds(start) || !walkable(g.Get(start)) {
		return nil
	}

	distances := make(map[Point]int)
	queue := []Point{start}
	distances[start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, next := range g.Neighbors4(current) {
			if _, seen := distances[next]; !seen && walkable(g.Get(next)) {
				distances[next] = distances[current] + 1
				queue = append(queue, next)
			}
		}
	}

	return distances
}

// PriorityQueue for A* pathfinding
type priorityQueueItem struct {
	point    Point
	priority int
	index    int
}

type priorityQueue []*priorityQueueItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*priorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// AStar performs A* pathfinding from start to goal
// cost function returns the cost to move to a given point
// Returns the path and total cost, or nil and -1 if no path exists
func AStar[T any](g *Grid[T], start, goal Point, cost func(Point) int, walkable func(T) bool) ([]Point, int) {
	if !g.InBounds(start) || !g.InBounds(goal) {
		return nil, -1
	}

	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &priorityQueueItem{point: start, priority: 0})

	cameFrom := make(map[Point]Point)
	costSoFar := make(map[Point]int)
	costSoFar[start] = 0

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*priorityQueueItem).point

		if current == goal {
			// Reconstruct path
			path := []Point{}
			for p := goal; p != start; p = cameFrom[p] {
				path = append([]Point{p}, path...)
			}
			path = append([]Point{start}, path...)
			return path, costSoFar[goal]
		}

		for _, next := range g.Neighbors4(current) {
			if !walkable(g.Get(next)) {
				continue
			}

			newCost := costSoFar[current] + cost(next)
			if oldCost, exists := costSoFar[next]; !exists || newCost < oldCost {
				costSoFar[next] = newCost
				priority := newCost + next.ManhattanDistance(goal)
				heap.Push(&pq, &priorityQueueItem{point: next, priority: priority})
				cameFrom[next] = current
			}
		}
	}

	return nil, -1
}
