#!/bin/bash
# Generate a new year's folder structure
# Usage: ./scripts/newyear.sh 2024

set -e

if [ -z "$1" ]; then
    echo "Usage: ./scripts/newyear.sh <year>"
    echo "Example: ./scripts/newyear.sh 2024"
    exit 1
fi

YEAR=$1

if [ -d "$YEAR" ]; then
    echo "Year $YEAR already exists"
    exit 1
fi

mkdir -p "$YEAR"

# Create register.go
cat > "$YEAR/register.go" << EOF
// Package y$YEAR registers all solutions for Advent of Code $YEAR.
package y$YEAR

import (
	"aoc/cmd/aoc"
	// Add day imports as you solve them:
	// "aoc/$YEAR/day01"
)

func init() {
	// Register solutions as you solve them:
	// main.Register($YEAR, 1, day01.Solution{})
	_ = main.Solutions // remove when you add first day
}
EOF

echo "✅ Created $YEAR/"
echo ""
echo "Next steps:"
echo "  1. Add import to cmd/aoc/main.go:"
echo "     _ \"aoc/$YEAR\""
echo ""
echo "  2. Generate first day:"
echo "     ./scripts/newday.sh $YEAR 1"
