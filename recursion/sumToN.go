package recursion

// SumToN возвращает сумму чисел от 1 до n.
// SumToN(5) = 1 + 2 + 3 + 4 + 5 = 15
// Для n <= 0 вернуть 0.
func SumToN(n int) int {
	// TODO
	if n <= 0 {
		return 0
	}
	return n + SumToN(n-1)
}
