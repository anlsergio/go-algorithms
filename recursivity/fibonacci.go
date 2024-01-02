package recursivity

// Fibonacci returns the nth number according to the
// Fibonacci's sequence, given the kickoff sequence:
//
//   - F(0) = 0
//   - F(1) = 1
func Fibonacci(nth uint) uint {
	if nth <= 1 {
		return nth
	}

	return Fibonacci(nth-1) + Fibonacci(nth-2)
}
