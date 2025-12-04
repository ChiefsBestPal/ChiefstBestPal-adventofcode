#!/bin/bash
# Generate a new day's files
# Usage: ./scripts/newday.sh 2025 5

set -e

if [ -z "$1" ] || [ -z "$2" ]; then
    echo "Usage: ./scripts/newday.sh <year> <day>"
    echo "Example: ./scripts/newday.sh 2025 5"
    exit 1
fi

YEAR=$1
DAY=$(printf "%02d" "$2")
DIR="$YEAR/day$DAY"

if [ ! -d "$YEAR" ]; then
    echo "Year $YEAR doesn't exist. Run ./scripts/newyear.sh $YEAR first."
    exit 1
fi

if [ -d "$DIR" ]; then
    echo "Directory $DIR already exists"
    exit 1
fi

mkdir -p "$DIR"

# Create solution file
cat > "$DIR/solution.go" << EOF
package day$DAY

import (
\	"aoc/shared/parser"
)

type Solution struct{}

func (Solution) Part1(input string) any {
	lines := parser.Lines(input)
	_ = lines // TODO
	return 0
}

func (Solution) Part2(input string) any {
	return 0
}
EOF

# Create test file
cat > "$DIR/solution_test.go" << EOF
package day$DAY

import "testing"

const example = \`\`

func TestPart1(t *testing.T) {
	got := Solution{}.Part1(example)
	want := 0
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Solution{}.Part2(example)
	want := 0
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution{}.Part1(example)
	}
}
EOF

# Create input files
touch "$DIR/input.txt"
touch "$DIR/example.txt"

echo "✅ Created $DIR/"
echo ""
echo "Next steps:"
echo "  1. Add to $YEAR/register.go:"
echo "     import \"aoc/$YEAR/day$DAY\""
echo "     aoc.Register($YEAR, $2, day$DAY.Solution{})"
echo ""
echo "  2. Paste example into $DIR/example.txt"
echo "  3. Paste input into $DIR/input.txt"
echo "  4. Solve!"
