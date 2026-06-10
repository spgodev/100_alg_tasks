package slice

// CountInRange возвращает количество элементов в диапазоне [left, right] включительно. Если left > right, возвращает 0.
func CountInRange(src []int, left int, right int) int {
	// TODO: реализовать функцию.
	if left > right {
		return 0
	}
	count := 0
	for i := range src {
		if src[i] >= left && src[i] <= right {
			count++
		}
	}
	return count
}
