package slice

// FirstPositiveIndex возвращает индекс первого положительного числа. Если такого числа нет, возвращает -1.
func FirstPositiveIndex(src []int) int {
	// TODO: реализовать функцию.
	result := -1
	for i := range src {
		if src[i] > 0 {
			result = i
			break
		}
	}
	return result
}
