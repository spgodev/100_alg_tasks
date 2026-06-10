package slice

// SumNegative возвращает сумму всех отрицательных чисел из src. Если отрицательных нет, возвращает 0.
func SumNegative(src []int) int {
	// TODO: реализовать функцию.
	sum := 0
	for _, value := range src {
		if value < 0 {
			sum += value
		}
	}
	return sum
}
