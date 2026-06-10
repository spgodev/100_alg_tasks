package slice

// DoubleSlice возвращает новый слайс, где каждый элемент src умножен на 2.
func DoubleSlice(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src))
	copy(result, src)
	for i := range result {
		result[i] = result[i] * 2
	}
	return result
}
