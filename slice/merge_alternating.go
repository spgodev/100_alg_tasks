package slice

// MergeAlternating возвращает новый слайс, чередуя элементы a и b. Остаток более длинного слайса дописывается в конец.
func MergeAlternating(a []int, b []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(a)+len(b))
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(a) {
			result = append(result, a[i])
		}
		if i < len(b) {
			result = append(result, b[i])
		}
	}

	return result
}
