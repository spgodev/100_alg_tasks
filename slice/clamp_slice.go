package slice

// ClampSlice возвращает новый слайс, где значения меньше min заменены на min, а больше max — на max.
func ClampSlice(src []int, min int, max int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src))
	copy(result, src)
	for i := range result {
		if result[i] < min {
			result[i] = min
		}
		if result[i] > max {
			result[i] = max
		}
	}
	return result
}
