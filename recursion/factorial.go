package recursion

// Factorial calculates the factorial of num.
func Factorial(num uint) uint {
	if num <= 1 {
		return 1
	}
	return num * Factorial(num-1)
}
