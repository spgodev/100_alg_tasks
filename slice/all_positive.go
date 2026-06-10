package slice

// AllPositive возвращает true, если все элементы src строго положительные. Для nil или пустого слайса возвращает false.
func AllPositive(src []int) bool {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return false
	}
	for i := range src {
		if src[i] <= 0 {
			return false
		}
	}
	return true
}
