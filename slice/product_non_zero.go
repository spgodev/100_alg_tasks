package slice

// ProductNonZero возвращает произведение всех ненулевых элементов. Если ненулевых элементов нет, возвращает 0.
func ProductNonZero(src []int) int {
	// TODO: реализовать функцию.
	result := 1
	foundNonZero := false
	if len(src) == 0 {
		return 0
	}
	for i := range src {
		if src[i] != 0 {
			foundNonZero = true
			result = result * src[i]
		}
	}
	if foundNonZero == false {
		return 0
	}
	return result
}
