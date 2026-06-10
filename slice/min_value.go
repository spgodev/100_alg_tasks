package slice

// MinValue возвращает минимальное значение в src. Для nil или пустого слайса возвращает 0.
func MinValue(src []int) int {
	// TODO: реализовать функцию.
	result := 0
	if len(src) <= 0 {
		return result
	}
	result = src[0]
	for i := range src {
		if src[i] < result {
			result = src[i]
		}
	}
	return result
}
