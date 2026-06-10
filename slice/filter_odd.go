package slice

// FilterOdd возвращает новый слайс только с нечётными числами в исходном порядке.
func FilterOdd(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for i := range src {
		if src[i]%2 == 1 || src[i]%2 == -1 {
			result = append(result, src[i])
		}
	}
	return result
}
