package slice

// FilterPositive возвращает новый слайс только со строго положительными числами.
func FilterPositive(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for i := range src {
		if src[i] > 0 {
			result = append(result, src[i])
		}
	}
	return result
}
