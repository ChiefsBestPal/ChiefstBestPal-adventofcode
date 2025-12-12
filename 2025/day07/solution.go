package day07

import (
	"aoc/shared/grid"
	"sync"
)

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
	g := Parse(input)
	start, found := g.Find(func(r rune) bool { return r == 'S' })
	if !found {
		return 0
	}

	var mu sync.Mutex
	paths := make(map[grid.Point]int)
	paths[grid.Point{X: start.X, Y: start.Y + 1}] = 1
	timelines := 0

	done := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	process := func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			default:
			}

			mu.Lock()
			if len(paths) == 0 {
				mu.Unlock()
				return
			}

			current := make(map[grid.Point]int)
			for k, v := range paths {
				current[k] = v
			}
			paths = make(map[grid.Point]int)
			mu.Unlock()

			next := make(map[grid.Point]int)
			exitCount := 0
			for p, count := range current {
				if !g.InBounds(p) {
					exitCount += count
					continue
				}

				cell := g.Get(p)
				if cell == '.' || cell == 'S' {
					next[grid.Point{X: p.X, Y: p.Y + 1}] += count
				} else if cell == '^' {
					next[grid.Point{X: p.X - 1, Y: p.Y + 1}] += count
					next[grid.Point{X: p.X + 1, Y: p.Y + 1}] += count
				}
			}

			mu.Lock()
			timelines += exitCount
			for p, count := range next {
				paths[p] += count
			}
			mu.Unlock()
		}
	}

	go process()
	go process()

	wg.Wait()
	close(done)

	return timelines
}
