package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	aocpkg "aoc/aoc"

	// Register years here (blank import to trigger init)
	_ "aoc/2025"
)

func main() {
	year := flag.Int("year", 2025, "Year (2015-2025)")
	day := flag.Int("day", 0, "Day to run (1-25)")
	example := flag.Bool("example", false, "Use example.txt instead of input.txt")
	all := flag.Bool("all", false, "Run all implemented days for the year")
	flag.Parse()

	if *all {
		runAll(*year)
		return
	}

	if *day < 1 || *day > 25 {
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd/aoc -year 2025 -day 1           Run day 1 of 2025")
		fmt.Println("  go run ./cmd/aoc -year 2025 -day 1 -example  Use example input")
		fmt.Println("  go run ./cmd/aoc -year 2025 -all             Run all days for 2025")
		os.Exit(1)
	}

	runDay(*year, *day, *example)
}

func runDay(year, day int, useExample bool) {
	solutions, ok := aocpkg.Solutions[year]
	if !ok {
		fmt.Printf("Year %d: not implemented\n", year)
		return
	}

	sol, ok := solutions[day]
	if !ok {
		fmt.Printf("Day %02d (%d): not implemented\n", day, year)
		return
	}

	input, err := loadInput(year, day, useExample)
	if err != nil {
		fmt.Printf("Day %02d (%d): %v\n", day, year, err)
		return
	}

	fmt.Printf("--- Day %02d (%d) ---\n", day, year)

	start := time.Now()
	p1 := sol.Part1(input)
	t1 := time.Since(start)

	start = time.Now()
	p2 := sol.Part2(input)
	t2 := time.Since(start)

	fmt.Printf("Part 1: %-20v (%v)\n", p1, t1.Round(time.Microsecond))
	fmt.Printf("Part 2: %-20v (%v)\n", p2, t2.Round(time.Microsecond))
}

func runAll(year int) {
	solutions, ok := aocpkg.Solutions[year]
	if !ok {
		fmt.Printf("Year %d: not implemented\n", year)
		return
	}

	fmt.Printf("Advent of Code %d\n\n", year)

	totalTime := time.Duration(0)
	stars := 0

	for day := 1; day <= 25; day++ {
		sol, ok := solutions[day]
		if !ok {
			continue
		}

		input, err := loadInput(year, day, false)
		if err != nil {
			fmt.Printf("Day %02d: %v\n", day, err)
			continue
		}

		start := time.Now()
		p1 := sol.Part1(input)
		p2 := sol.Part2(input)
		elapsed := time.Since(start)
		totalTime += elapsed

		dayStars := ""
		if p1 != nil && p1 != 0 && p1 != "" {
			dayStars += "★"
			stars++
		}
		if p2 != nil && p2 != 0 && p2 != "" {
			dayStars += "★"
			stars++
		}

		fmt.Printf("Day %02d: %-4s Part1=%-14v Part2=%-14v (%v)\n",
			day, dayStars, p1, p2, elapsed.Round(time.Microsecond))
	}

	fmt.Println()
	fmt.Printf("Total: %d stars in %v\n", stars, totalTime.Round(time.Microsecond))
}

func loadInput(year, day int, useExample bool) (string, error) {
	filename := "input.txt"
	if useExample {
		filename = "example.txt"
	}

	path := filepath.Join(fmt.Sprintf("%d", year), fmt.Sprintf("day%02d", day), filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read %s", path)
	}

	return string(data), nil
}
