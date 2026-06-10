package slice

// CountGreaterThan возвращает количество элементов src, которые строго больше limit.
func CountGreaterThan(src []int, limit int) int {
	// TODO: реализовать функцию.
	counter := 0
	for i := range src {
		if src[i] > limit {
			counter++
		}
	}
	return counter
}
