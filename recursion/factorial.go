package recursion

// Factorial возвращает факториал числа n.
// Factorial(0) = 1
// Factorial(5) = 120
func Factorial(n int) int {
	// TODO
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
