package math

// Mod returns the mathematical modulo (always non-negative)
// Unlike Go's % operator which can return negative values
func Mod(a, m int) int {
	return ((a % m) + m) % m
}

// Abs returns the absolute value of x
func Abs[T ~int | ~int64 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the minimum of two values
func Min[T ~int | ~int64 | ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two values
func Max[T ~int | ~int64 | ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// GCD returns the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
