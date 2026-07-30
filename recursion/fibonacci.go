package recursion

// Fibonacci возвращает n-е число Фибоначчи.
// Fibonacci(0) = 0
// Fibonacci(1) = 1
// Fibonacci(6) = 8
func Fibonacci(n int) int {
	// TODO
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
