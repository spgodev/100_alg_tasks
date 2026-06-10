package slice

// LastIndex возвращает индекс последнего вхождения target в src. Если target не найден, возвращает -1.
func LastIndex(src []int, target int) int {
	// TODO: реализовать функцию.
	result := -1
	for i := range src {
		if src[i] == target {
			result = i
		}
	}
	return result
}
