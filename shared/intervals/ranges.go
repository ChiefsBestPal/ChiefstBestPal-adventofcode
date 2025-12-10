package interval

import (
	"slices"
)

type Range struct {
	Lo, Hi int
}

func Sort(ranges []Range) {
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Lo != b.Lo {
			return a.Lo - b.Lo
		}
		return a.Hi - b.Hi
	})
}

func (r Range) Contains(val int) bool {
	return val >= r.Lo && val <= r.Hi
}

func (r Range) Size() int {
	return r.Hi - r.Lo + 1
}

func (a Range) IsOverlapping(b Range) bool {
	return a.Lo <= b.Hi && b.Lo <= a.Hi
}

func (a Range) IsAdjacent(b Range) bool {
	return a.Hi+1 == b.Lo || b.Hi+1 == a.Lo
}

func (a Range) CanMerge(b Range) bool {
	return a.IsOverlapping(b) || a.IsAdjacent(b)
}

// Merge merges all overlapping or adjacent ranges in a slice.
// Input slice assumed to be sorted
func Merge(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	} else if len(ranges) == 2 {
		var a = ranges[0]
		var b = ranges[1]
		if !a.CanMerge(b) {
			return ranges
		}
		return []Range{
			Range{
				Lo: min(a.Lo, b.Lo),
				Hi: max(a.Hi, b.Hi)},
		}
	}

	merged := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		cur := ranges[i]

		if last.CanMerge(cur) {
			// merge directly into "last"
			if cur.Hi > last.Hi {
				last.Hi = cur.Hi
			}
		} else {
			merged = append(merged, cur)
		}
	}

	return merged
}
func MergeUnsorted(ranges []Range) []Range {
	Sort(ranges)
	return Merge(ranges)
}
