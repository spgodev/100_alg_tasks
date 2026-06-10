package slice

// CountLessThan возвращает количество элементов src, которые строго меньше limit.
func CountLessThan(src []int, limit int) int {
	// TODO: реализовать функцию.
	count := 0
	for i := range src {
		if src[i] < limit {
			count++
		}
	}
	return count
}
