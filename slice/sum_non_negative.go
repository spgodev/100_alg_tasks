package slice

// SumNonNegative возвращает сумму чисел, которые больше либо равны нулю.
func SumNonNegative(src []int) int {
	// TODO: реализовать функцию.
	sum := 0
	for _, value := range src {
		if value >= 0 {
			sum += value
		}
	}
	return sum
}
