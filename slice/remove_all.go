package slice

// RemoveAll возвращает новый слайс без всех элементов, равных target. Порядок остальных элементов сохраняется.
func RemoveAll(src []int, target int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for i := range src {
		if src[i] != target {
			result = append(result, src[i])
		}
	}
	return result
}
