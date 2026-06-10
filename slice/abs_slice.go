package slice

// AbsSlice возвращает новый слайс из модулей элементов src.
func AbsSlice(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src))
	copy(result, src)
	for i := range result {
		if result[i] < 0 {
			result[i] = result[i] * -1
		}
	}
	return result
}
