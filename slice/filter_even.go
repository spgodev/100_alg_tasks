package slice

// FilterEven возвращает новый слайс только с чётными числами в исходном порядке.
func FilterEven(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for i := range src {
		if src[i]%2 == 0 {
			result = append(result, src[i])
		}
	}
	return result
}
