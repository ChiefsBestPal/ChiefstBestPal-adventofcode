# Advent of Code

My Go solutions for [Advent of Code](https://adventofcode.com/) to share with friends and coworkers.

## Project Structure

```
aoc/
├── cmd/aoc/main.go           # CLI entry point
├── 2025/                     # Year folder
│   ├── day01/
│   │   ├── day01.go          # Solution
│   │   ├── day01_test.go     # Tests + benchmarks
│   │   ├── input.txt         # Puzzle input (gitignored)
│   │   └── example.txt       # Input from public AOC example 
│   ├── day02/
│   │   └── ...
├── shared/                   # Common utilities (year/day-agnostic)
│   └── ...
├── scripts/
│   └── ...
├── go.mod
├── go.sum (optional, for max reproducibility only)
└── README.md
```

## Progress

### 2025
| Day | Stars | Notes |
|-----|:-----:|-------|
| 01  | ... | Safe dial |
| 02  | | |
| ... | | |
## Commands

```bash
# Run specific day
go run ./cmd/aoc -year 2025 -day 1

# Run with example input
go run ./cmd/aoc -year 2025 -day 5 -example

# Run all days for a year
go run ./cmd/aoc -year 2025 -all

# Run tests for a year
go test ./2025/...

# Run tests for a specific day
go test ./2025/day01 -v

# All tests
go test ./...

# Benchmarks
go test ./2025/day01 -bench=. -benchmem

# Generate new day
./scripts/newday.sh 2025 5
```